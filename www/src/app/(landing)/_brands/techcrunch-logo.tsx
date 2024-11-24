import * as React from 'react';
import { CustomSVGProps } from '../_types';

const TechCrunchSVG: React.FC<CustomSVGProps> = (props) => (
  <svg
    width={49}
    height={24}
    viewBox='0 0 49 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g clipPath='url(#clip0_406_79)'>
      <path
        d='M24.5716 0.0322266V8.07547H16.5287V23.9998H8.3885V8.07547H0.280762V0.0322266H24.5716Z'
        fill={props.fillcolor}
      />
      <path d='M48.6028 0H32.6792V8.10811H48.6028V0Z' fill={props.fillcolor} />
      <path
        d='M48.6031 16.2165V24.0003H24.5718V8.1084H32.712V16.2165H48.6031Z'
        fill={props.fillcolor}
      />
    </g>
    <defs>
      <clipPath id='clip0_406_79'>
        <rect
          width={48.3221}
          height={24}
          fill='white'
          transform='translate(0.280762)'
        />
      </clipPath>
    </defs>
  </svg>
);
export default TechCrunchSVG;
