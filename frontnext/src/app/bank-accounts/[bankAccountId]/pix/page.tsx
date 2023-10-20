import { Box } from '@mui/material';
import Grid2 from '@mui/material/Unstable_Grid2/Grid2';
import { PixKeyList } from './PixKeyList';
import { PixKeyForm } from './PixKeyForm';

export default function PixDashboardPage({
  params,
}: {
  params: { bankAccountId: string };
}) {
  return (
    <Box>
      <Grid2 container spacing={8}>
        <Grid2 xs={12} sm={6}>
          <PixKeyForm bankAccountId={params.bankAccountId} />
        </Grid2>

        <Grid2 xs={12} sm={6}>
          <div>
            <PixKeyList bankAccountId={params.bankAccountId} />
          </div>
        </Grid2>
      </Grid2>
    </Box>
  );
}
