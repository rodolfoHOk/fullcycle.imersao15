import {
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Typography,
} from '@mui/material';
import { StarBorder } from '@mui/icons-material';
import { PixKey } from '@/models/models';

export async function getPixKeys(bankAccountId: string): Promise<PixKey[]> {
  const nestApiUrl = process.env.NEXT_PUBLIC_NEST_API_URL;
  const response = await fetch(
    `${nestApiUrl}/bank-accounts/${bankAccountId}/pix-keys`,
    {
      next: {
        tags: [`pix-keys-${bankAccountId}`],
      },
    }
  ); //cache sob demanda

  return response.json();
}

type PixKeyListProps = {
  bankAccountId: string;
};

export async function PixKeyList(props: PixKeyListProps) {
  const pixKeys = await getPixKeys(props.bankAccountId);

  return (
    <div>
      <Typography variant="h5">Minhas chaves pix</Typography>

      <List>
        {pixKeys.map((pixKey) => (
          <ListItem disablePadding key={pixKey.id}>
            <ListItemButton>
              <ListItemIcon>
                <StarBorder />
              </ListItemIcon>

              <ListItemText primary={pixKey.key} secondary={pixKey.kind} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </div>
  );
}
