import { CardAction } from '@/components/CardAction/CardAction';
import { BankAccount } from '@/models/models';
import { Typography } from '@mui/material';
import Grid2 from '@mui/material/Unstable_Grid2/Grid2';

async function getBankAccounts(): Promise<BankAccount[]> {
  const nestApiUrl = process.env.NEXT_PUBLIC_NEST_API_URL;
  const response = await fetch(`${nestApiUrl}/bank-accounts`, {
    next: {
      revalidate: 10,
    },
  });

  return response.json();
}

export default async function BankAccountsPage() {
  const bankAccounts = await getBankAccounts();

  return (
    <div>
      <Typography variant="h4">Contas bancárias</Typography>

      <Grid2 container spacing={2} mt={1}>
        {bankAccounts.map((bankAccount) => (
          <Grid2 key={bankAccount.id} xs={12} sm={6} md={4}>
            <CardAction>
              <Typography variant="h5" component="div">
                {bankAccount.owner_name}
              </Typography>

              <Typography component="span">
                Conta: {bankAccount.account_number}
              </Typography>
            </CardAction>
          </Grid2>
        ))}
      </Grid2>
    </div>
  );
}
