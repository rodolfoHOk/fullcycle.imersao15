import { Box } from '@mui/material';
import { WithdrawForm } from './WithDrawForm';

export default async function WithdrawPage({
  params,
}: {
  params: { bankAccountId: string };
}) {
  return (
    <Box>
      <WithdrawForm bankAccountId={params.bankAccountId} />
    </Box>
  );
}
