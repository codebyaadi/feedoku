import type { Metadata } from 'next';
import { geistMono, geistSans, outfit } from '@/app/fonts';
import { siteConfig } from '@/config/site';
import './globals.css';

export const metadata: Metadata = {
  title: {
    default: siteConfig.name,
    template: `%s - ${siteConfig.name}`,
  },
  description: siteConfig.description,
  keywords: [
    'golang',
    'nextjs',
    'rss feed aggregator',
    'golang rss feed',
    'golang rss aggregator',
    'golang rss',
    'react',
    'panda css',
  ],
  authors: [
    {
      name: 'codebyaadi',
      url: 'https://codebyaadi.netlify.app',
    },
  ],
  creator: 'codebyaadi',
  openGraph: {
    type: 'website',
    locale: 'en_IN',
    title: siteConfig.name,
    description: siteConfig.description,
    siteName: siteConfig.name,
    url: siteConfig.url,
    images: [
      {
        url: `${siteConfig.url}/og.jpg`,
        width: 1200,
        height: 630,
        alt: siteConfig.name,
      },
    ],
  },
  twitter: {
    card: 'summary_large_image',
    title: siteConfig.name,
    description: siteConfig.description,
    creator: '@codebyaadi',
    images: [
      {
        url: `${siteConfig.url}/og.jpg`,
        alt: siteConfig.name,
      },
    ],
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html
      lang='en'
      className={`${geistSans.variable} ${geistMono.variable} ${outfit.variable}`}
    >
      <body>{children}</body>
    </html>
  );
}
