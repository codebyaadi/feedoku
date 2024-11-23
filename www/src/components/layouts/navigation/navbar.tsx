import { Button } from '@/components/ui/button';
import { siteConfig } from '@/config/site';
import Link from 'next/link';
import React from 'react';

const Navbar = () => {
  return (
    <nav className='fixed top-0 z-[1] w-full bg-transparent px-0 font-outfit lg:px-6'>
      <div className='mx-4 flex items-center justify-between py-2'>
        <Link href='/' className='font-semibold text-amber-300'>
          {siteConfig.name}
        </Link>
        <div className='flex items-center justify-center gap-2 font-normal'>
          <Button
            variant='secondary'
            size='sm'
            className='rounded-sm shadow-none'
          >
            Log In
          </Button>
          <Button variant='amber' size='sm'>
            Sign Up
          </Button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
