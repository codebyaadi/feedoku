import React from 'react';
import Link from 'next/link';
import { ModeToggle } from '@/components/mode-toggle';
import { Button } from '@/components/ui/button';
import { siteConfig } from '@/config/site';

const Navbar = () => {
  return (
    <nav className='fixed top-0 z-[1] w-full bg-transparent px-0 font-outfit lg:px-6'>
      <div className='mx-4 flex items-center justify-between py-2'>
        <Link href='/' className='font-semibold text-amber-300'>
          {siteConfig.name}
        </Link>
        <div className='flex items-center justify-center gap-2 font-normal'>
          <ModeToggle />
          <Button
            variant='secondary'
            size='sm'
            className='rounded-sm shadow-none'
          >
            <Link href='/signin'>Log In</Link>
          </Button>
          <Button variant='amber' size='sm'>
            <Link href='/signup'>Sign Up</Link>
          </Button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
