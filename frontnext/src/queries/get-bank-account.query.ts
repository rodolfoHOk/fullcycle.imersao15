import { BankAccount } from '@/models/models';

export async function getBankAccount(
  bankAccountId: string
): Promise<BankAccount> {
  const nestApiUrl = process.env.NEXT_PUBLIC_NEST_API_URL;
  const response = await fetch(`${nestApiUrl}/bank-accounts/${bankAccountId}`, {
    next: {
      revalidate: 20,
    },
  });

  return response.json();
}
