import { BankAccount } from '@/models/models';

async function getBankAccounts(): Promise<BankAccount[]> {
  const nestApiUrl = process.env.NEXT_PUBLIC_NEST_API_URL;
  const response = await fetch(`${nestApiUrl}/bank-accounts`, {
    next: {
      revalidate: 10,
    },
  });

  return response.json();
}

export default function BankAccountsPage() {
  return (
    <div>
      <h1>BankAccountsPage</h1>
    </div>
  );
}
