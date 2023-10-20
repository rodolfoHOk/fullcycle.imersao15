import { Typography } from '@mui/material';
import Grid2 from '@mui/material/Unstable_Grid2/Grid2';
import { CurrentBalance } from '@/components/CurrentBalance/CurrentBalance';
import { CardAction } from '@/components/CardAction/CardAction';
import { LatestTransactions } from './LatestTransactions';
import { Transaction } from '@/models/models';

async function getTransactions(bankAccountId: string): Promise<Transaction[]> {
  const nestApiUrl = process.env.NEXT_PUBLIC_NEST_API_URL;
  const response = await fetch(
    `${nestApiUrl}/bank-accounts/${bankAccountId}/transactions`,
    {
      next: {
        revalidate: 10,
      },
    }
  );

  return response.json();
}

export default async function BankAccountDashboardPage({
  params,
}: {
  params: { bankAccountId: string };
}) {
  const transactions = await getTransactions(params.bankAccountId);

  return (
    <Grid2 container spacing={2}>
      <Grid2 xs={12} lg={6}>
        <div>
          <CurrentBalance bankAccountId={params.bankAccountId} />
        </div>
      </Grid2>

      <Grid2 container xs={12} lg={6} spacing={1}>
        <Grid2 xs={6}>
          <CardAction sx={{ display: 'flex', alignItems: 'center' }}>
            <Typography component="span" color={'primary'}>
              Transferência
            </Typography>
          </CardAction>
        </Grid2>

        <Grid2 xs={6}>
          <CardAction sx={{ display: 'flex', alignItems: 'center' }}>
            <Typography component="span" color={'primary'}>
              Nova chave pix
            </Typography>
          </CardAction>
        </Grid2>
      </Grid2>

      <Grid2 xs={12}>
        <Typography variant="h5">Últimos lançamentos</Typography>

        <LatestTransactions transactions={transactions} />
      </Grid2>
    </Grid2>
  );
}
