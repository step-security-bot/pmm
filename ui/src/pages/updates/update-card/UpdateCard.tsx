import {
  Button,
  Card,
  CardActions,
  CardContent,
  Stack,
  Link,
  Typography,
  Skeleton,
  Alert,
} from '@mui/material';
import { FC, useMemo, useState } from 'react';
import { formatTimestamp } from 'utils/formatTimestamp';
import { PMM_HOME_URL } from 'constants';
import { Messages } from './UpdateCard.messages';
import { FetchingIcon } from 'components/fetching-icon';
import { useCheckUpdates, useStartUpdate } from 'hooks/api/useUpdates';
import { formatVersion } from './UpdateCard.utils';
import { enqueueSnackbar } from 'notistack';
import { useWaitForReadiness } from 'hooks/api/useReadiness';
import { UpdateStatus } from 'types/updates.types';
import { KeyboardDoubleArrowUp } from '@mui/icons-material';
import { UpdateInfo } from '../update-info';
import { UpdateInProgressCard } from '../update-in-progress-card';

export const UpdateCard: FC = () => {
  const { isLoading, data, error, isRefetching, refetch } = useCheckUpdates();
  const { mutate: startUpdate } = useStartUpdate();
  const { waitForReadiness } = useWaitForReadiness();
  const [status, setStatus] = useState<UpdateStatus>(UpdateStatus.Pending);
  const isUpToDate = useMemo(
    () => data?.installed.version === data?.latest?.version,
    [data]
  );
  const updateInProgress = useMemo(
    () =>
      status === UpdateStatus.Updating ||
      status === UpdateStatus.Restarting ||
      status === UpdateStatus.Completed,
    [status]
  );

  const handleStartUpdate = async () => {
    setStatus(UpdateStatus.Updating);
    startUpdate(
      {},
      {
        onSuccess: async () => {
          setStatus(UpdateStatus.Restarting);

          await waitForReadiness();

          setStatus(UpdateStatus.Completed);
        },
        onError: () => {
          setStatus(UpdateStatus.Error);
          enqueueSnackbar(Messages.error, {
            variant: 'error',
          });
        },
      }
    );
  };

  if (isLoading)
    return (
      <Card>
        <CardContent>
          <Stack spacing={1}>
            <Skeleton />
            <Skeleton />
          </Stack>
        </CardContent>
      </Card>
    );

  if (!data || error) {
    return (
      <Card>
        <CardContent>
          <Alert severity="error">{Messages.fetchError}</Alert>
        </CardContent>
      </Card>
    );
  }

  if (updateInProgress && data.latest) {
    return <UpdateInProgressCard versionInfo={data.latest} status={status} />;
  }

  return (
    <Card sx={{ p: 1 }}>
      <CardContent>
        {isUpToDate && (
          <Alert
            severity="success"
            sx={{
              mb: 2,
            }}
          >
            {Messages.upToDate}
          </Alert>
        )}
        <Stack spacing={1}>
          {data.updateAvailable && data?.latest?.version && (
            <Typography variant="h4">
              {Messages.newUpdateAvailable(data.latest.version)}
            </Typography>
          )}
          <Typography>
            <Typography fontWeight="bold" component="strong">
              {Messages.runningVersion}
            </Typography>
            {data?.installed && formatVersion(data.installed)}
          </Typography>
          {data.updateAvailable && data.latest && (
            <Typography>
              <Typography fontWeight="bold" component="strong">
                {Messages.newVersion}
              </Typography>
              {data?.latest && formatVersion(data.latest)}
            </Typography>
          )}
          {data.lastCheck && (
            <Typography>
              <Typography fontWeight="bold" component="strong">
                {Messages.lastChecked}
              </Typography>{' '}
              {formatTimestamp(data?.lastCheck)}
            </Typography>
          )}
        </Stack>
        {data.updateAvailable && <UpdateInfo />}
      </CardContent>
      {data.updateAvailable ? (
        <CardActions>
          <Button
            endIcon={<KeyboardDoubleArrowUp />}
            variant="contained"
            onClick={handleStartUpdate}
          >
            {Messages.updateNow}
          </Button>
        </CardActions>
      ) : (
        <CardActions>
          <Button
            startIcon={<FetchingIcon isFetching={isRefetching} />}
            variant="contained"
            onClick={() => refetch()}
          >
            {isRefetching ? Messages.checking : Messages.checkNow}
          </Button>
          <Link href={PMM_HOME_URL}>
            <Button variant="outlined">{Messages.home}</Button>
          </Link>
        </CardActions>
      )}
    </Card>
  );
};