'use client';
import { AnimatedTooltip } from '@/components/_ui/animated-tooltip';
import React from 'react';
const brand = [
  {
    id: 1,
    name: 'Reddit',
    designation: 'Social Platform',
    image:
      'https://res.cloudinary.com/dqcejxdbf/image/upload/t_512by512/v1732293552/reddit_ixll7v.png',
  },
  {
    id: 2,
    name: 'Dev.to',
    designation: 'Developer Community',
    image:
      'https://res.cloudinary.com/dqcejxdbf/image/upload/v1732373410/dev-black_vtkiqc.png',
  },
  {
    id: 3,
    name: 'Hashnode',
    designation: 'Tech Blogging',
    image:
      'https://res.cloudinary.com/dqcejxdbf/image/upload/v1732294636/hashnode_efrt.png',
  },
  {
    id: 4,
    name: 'FreeCodeCamp',
    designation: 'Learn Coding',
    image:
      'https://res.cloudinary.com/dqcejxdbf/image/upload/v1732373275/freecodecamp_pnhouc.png',
  },
  {
    id: 5,
    name: 'MIT News',
    designation: 'Research Insights',
    image:
      'https://res.cloudinary.com/dqcejxdbf/image/upload/v1732373638/logo-colors-mit-red_zuk0qh.png',
  },
  {
    id: 6,
    name: 'Stack Overflow',
    designation: 'Developer Community',
    image:
      'https://res.cloudinary.com/dqcejxdbf/image/upload/v1732293969/stackoverflow-logo_bluxw2.png',
  },
];

export function AvatarStack() {
  return (
    <div className='mb-1 flex w-full flex-row items-center justify-center'>
      <AnimatedTooltip items={brand} />
    </div>
  );
}
