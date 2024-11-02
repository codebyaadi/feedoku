import Navbar from '@/components/layouts/navigation/navbar';
import { css } from 'styled-system/css';

export default function LandingLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div
      className={css({
        position: 'relative',
      })}
    >
      <Navbar />
      <div
        className={css({
          position: 'absolute',
          top: '0',
          h: '1/3',
          w: 'full',
          bgGradient: 'to-b',
          gradientFrom: 'accent.a3',
          gradientTo: 'transparent'
        })}
      >

      </div>
      {children}
    </div>
  );
}
