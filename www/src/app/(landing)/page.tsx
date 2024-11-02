import { ArrowRightIcon } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { css } from 'styled-system/css';
import { flex } from 'styled-system/patterns';
import AvatarStack from './_components/avatar-stack';
import { avatars } from './_components/constant';
import { RatingGroup } from '@/components/ui/rating-group';
import Brands from './_components/brands';

export default function Home() {
  return (
    <div className={css({ fontFamily: 'body' })}>
      <main
        className={flex({
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          mx: { base: '4', lg: 'auto' },
          pt: { base: '32', lg: '48' },
        })}
      >
        <div
          className={flex({
            flexDirection: 'column',
            justifyContent: 'center',
            alignItems: 'center',
            maxW: { base: '4xl', lg: '3xl' },
          })}
        >
          <div
            className={css({
              display: 'flex',
              alignItems: 'center',
              flexDirection: {base: 'column', md: 'row', lg: 'row'},
              gap: {base: '2', md: '4', lg: '4'},
            })}
          >
            <AvatarStack avatars={avatars} maxVisible={4} />
            <RatingGroup
              value={5}
              className={css({
                display: 'flex',
                justifyContent: 'center',
                alignContent: 'center',
                gap: '0',
              })}
            >
              10+ feeds to subscribe
            </RatingGroup>
          </div>
          <div
            className={css({
              width: 'full',
              mb: { base: '10', lg: '16' },
              mt: { base: '2', lg: '8' },
              textAlign: 'center',
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
              Stay Informed. Stay Connected
            </h1>
            <p className={css({ color: 'sage.a11', lineHeight: 'tight' })}>
              Feedoku brings the power of real-time news and content aggregation
              to your fingertips. Discover, organize, and read your favorite RSS
              feeds from across the web in one simple platform.
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
        </div>
        <Brands />
      </main>
      <footer className=''></footer>
    </div>
  );
}
