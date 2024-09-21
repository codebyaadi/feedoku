import { ArrowRightIcon } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { css } from 'styled-system/css';
import { grid } from 'styled-system/patterns';

export default function Home() {
  return (
    <div className={css({ fontFamily: 'body' })}>
      <main
        className={grid({
          columns: { base: 1, lg: 2 },
          maxW: '7xl',
          mx: { base: '4', lg: 'auto' },
        })}
      >
        <div
          className={css({
            width: 'full',
            my: { base: '16', lg: '24' },
          })}
        >
          <h1
            className={css({
              fontSize: { base: '4xl', lg: '5xl' },
              fontWeight: 'semibold',
              color: 'accent.emphasized',
              lineHeight: 'none',
              mb: '3',
            })}
          >
            Stay Informed <br /> Stay Connected
          </h1>
          <p className={css({ color: 'sage.text' })}>
            Feedly brings the power of real-time news and content aggregation to
            your fingertips. Discover, organize, and read your favorite RSS
            feeds from across the web in one simple platform. Whether it's tech
            updates, global news, or niche blogs, Feedly keeps you connected to
            what matters.
          </p>
          <div
            className={css({
              mt: '4',
            })}
          >
            <Button
              className={css({
                mr: '2',
                fontWeight: 'medium',
              })}
            >
              Get Started
              <ArrowRightIcon />
            </Button>
            <Button variant='outline'>Learn more</Button>
          </div>
        </div>
      </main>
      <footer className=''></footer>
    </div>
  );
}
