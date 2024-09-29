'use client';

import Link from 'next/link';
import React from 'react';
import { css } from 'styled-system/css';
import { Button } from '@/components/ui/button';
import { useRouter } from 'next/navigation';
import { siteConfig } from '@/config/site';

const Navbar = () => {
  const router = useRouter();

  return (
    <nav
      className={css({
        position: 'absolute',
        w: 'full',
      })}
    >
      <div
        className={css({
          fontFamily: 'body',
          maxW: '7xl',
          mx: { base: '4', lg: 'auto' },
          py: '3',
          display: 'flex',
          flexDir: 'row',
          justifyContent: 'space-between',
          alignItems: 'center',
        })}
      >
        <div>
          <Link
            href='/'
            className={css({
              fontSize: 'larger',
              fontWeight: 'medium',
              color: 'accent.emphasized',
            })}
          >
            {siteConfig.name}
          </Link>
        </div>
        <div>
          <Button
            variant='ghost'
            className={css({
              mr: '2',
              fontWeight: 'medium',
            })}
          >
            Log In
          </Button>
          <Button
            className={css({
              fontWeight: 'medium',
            })}
            onClick={() => router.push('/signup')}
          >
            Sign Up
          </Button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
