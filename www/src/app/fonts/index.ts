import localFont from 'next/font/local';
import { Outfit, Roboto_Mono } from 'next/font/google';

export const outfit = Outfit({ subsets: ['latin'], variable: '--font-outfit' });
export const roboto = Roboto_Mono({
  subsets: ['latin'],
  variable: '--font-roboto-mono',
});

export const geistSans = localFont({
  src: './GeistVF.woff',
  variable: '--font-geist-sans',
  weight: '100 900',
});
export const geistMono = localFont({
  src: './GeistMonoVF.woff',
  variable: '--font-geist-mono',
  weight: '100 900',
});
