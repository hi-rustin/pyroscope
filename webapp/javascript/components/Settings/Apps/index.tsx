import React, { useEffect, useState } from 'react';
import cl from 'classnames';
import { useAppDispatch, useAppSelector } from '@webapp/redux/hooks';
import {
  selectApps,
  reloadApps,
  deleteApp,
  selectIsLoadingApps,
} from '@webapp/redux/reducers/settings';
import { addNotification } from '@webapp/redux/reducers/notifications';
import { type App } from '@webapp/models/app';
import TableUI from '@webapp/ui/Table';
import LoadingSpinner from '@webapp/ui/LoadingSpinner';
import TextField from '@webapp/ui/Form/TextField';
import { getAppTableRows } from './getAppTableRows';

import appsStyles from './Apps.module.css';
import tableStyles from '../SettingsTable.module.scss';

const headRow = [
  { name: '', label: 'Name', sortable: 0 },
  { name: '', label: '', sortable: 0 },
];

function Apps() {
  const dispatch = useAppDispatch();
  const apps = useAppSelector(selectApps);
  const isLoading = useAppSelector(selectIsLoadingApps);
  const [search, setSearchField] = useState('');
  const [appsInProcessing, setAppsInProcessing] = useState<string[]>([]);
  const [deletedApps, setDeletedApps] = useState<string[]>([]);

  useEffect(() => {
    dispatch(reloadApps());
  }, []);

  const displayApps =
    (apps &&
      apps.filter(
        (x) =>
          x.name.toLowerCase().indexOf(search.toLowerCase()) !== -1 &&
          !deletedApps.includes(x.name)
      )) ||
    [];

  const handleDeleteApp = (app: App) => {
    setAppsInProcessing([...appsInProcessing, app.name]);
    dispatch(deleteApp(app))
      .unwrap()
      .then(() => {
        setAppsInProcessing(appsInProcessing.filter((x) => x !== app.name));
        setDeletedApps([...deletedApps, app.name]);
        dispatch(
          addNotification({
            type: 'success',
            title: 'App has been deleted',
            message: `App ${app.name} has been successfully deleted`,
          })
        );
      })
      .catch(() => {
        setDeletedApps(deletedApps.filter((x) => x !== app.name));
        setAppsInProcessing(appsInProcessing.filter((x) => x !== app.name));
      });
  };

  const tableBodyProps =
    displayApps.length > 0
      ? {
          bodyRows: getAppTableRows(
            displayApps,
            appsInProcessing,
            handleDeleteApp
          ),
          type: 'filled' as const,
        }
      : {
          type: 'not-filled' as const,
          value: 'The list is empty',
          bodyClassName: appsStyles.appsTableEmptyMessage,
        };

  return (
    <>
      <h2 className={appsStyles.tabNameContrainer}>
        Apps
        {isLoading && !!apps ? <LoadingSpinner /> : null}
      </h2>
      <div className={appsStyles.searchContainer}>
        <TextField
          type="text"
          label="Search app"
          placeholder="Search app"
          variant="light"
          value={search}
          onChange={(v) => setSearchField(v.target.value)}
          name="Search app input"
        />
      </div>
      <TableUI
        className={cl(appsStyles.appsTable, tableStyles.settingsTable)}
        table={{ headRow, ...tableBodyProps }}
        isLoading={isLoading && !apps}
      />
    </>
  );
}

export default Apps;
