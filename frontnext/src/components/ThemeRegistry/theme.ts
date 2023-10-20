import { createTheme } from '@mui/material';
import { blue } from '@mui/material/colors';
import { Roboto } from 'next/font/google';

const roboto = Roboto({
  weight: ['300', '400', '500', '700'],
  subsets: ['latin'],
  display: 'swap',
});

const theme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: blue[500],
      light: blue[300],
      dark: blue[700],
      contrastText: '#fff',
    },
  },
  typography: {
    fontFamily: roboto.style.fontFamily,
  },
});

export default theme;
