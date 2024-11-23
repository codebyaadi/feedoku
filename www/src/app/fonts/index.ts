import { Outfit, Roboto_Mono } from 'next/font/google';

export const outfit = Outfit({ subsets: ['latin'], variable: '--font-outfit' });
export const roboto = Roboto_Mono({
  subsets: ['latin'],
  variable: '--font-roboto-mono',
});
