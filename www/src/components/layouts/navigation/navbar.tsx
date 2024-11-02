'use client';

import Link from 'next/link';
import React from 'react';
import { css } from 'styled-system/css';
import { Button } from '@/components/ui/button';
import { useRouter } from 'next/navigation';
import { siteConfig } from '@/config/site';
import { flex } from 'styled-system/patterns';

const Navbar = () => {
  const router = useRouter();

  return (
    <nav
      className={css({
        position: 'fixed',
        top: '0',
        zIndex: '1',
        w: 'full',
        // bgColor: 'bg.canvas',
        // borderWidth: '1',
        borderStyle: 'solid'
      })}
    >
      <div
        className={flex({
          fontFamily: 'body',
          mx: '4',
          py: '2',
          justifyContent: 'space-between',
          alignItems: 'center',
        })}
      >
        <div>
          <Link
            href='/'
            className={css({
              fontSize: 'lg',
              fontWeight: 'medium',
              color: 'accent.emphasized',
            })}
          >
            {siteConfig.name}
          </Link>
        </div>
        <div>
          <Button
            size='xs'
            variant='ghost'
            className={css({
              mr: '2',
              fontWeight: 'medium',
            })}
          >
            Log In
          </Button>
          <Button
            size='xs'
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
