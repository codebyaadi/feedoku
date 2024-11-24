import * as React from 'react';
import { CustomSVGProps } from '../_types';

const MediumSVG: React.FC<CustomSVGProps> = (props) => (
  <svg
    width={157}
    height={24}
    viewBox='0 0 157 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g clipPath='url(#clip0_406_73)'>
      <path
        d='M23.6289 12.1259C23.6289 18.5689 18.391 23.7923 11.9295 23.7923C5.46799 23.7923 0.230469 18.5705 0.230469 12.1259C0.230469 5.68134 5.46839 0.459961 11.9295 0.459961C18.3906 0.459961 23.6289 5.68291 23.6289 12.1259Z'
        fill={props.fillcolor}
      />
      <path
        d='M36.4631 12.1263C36.4631 18.1916 33.8439 23.1081 30.6134 23.1081C27.3828 23.1081 24.7637 18.19 24.7637 12.1263C24.7637 6.06257 27.3828 1.14453 30.6134 1.14453C33.8439 1.14453 36.4631 6.06257 36.4631 12.1263Z'
        fill={props.fillcolor}
      />
      <path
        d='M41.7125 12.1263C41.7125 17.5606 40.7913 21.9655 39.6551 21.9655C38.5189 21.9655 37.5977 17.559 37.5977 12.1263C37.5977 6.69367 38.5189 2.28711 39.6555 2.28711C40.792 2.28711 41.7125 6.69249 41.7125 12.1263Z'
        fill={props.fillcolor}
      />
      <path
        d='M74.0012 1.47175L74.0338 1.46466V1.21483H67.3804L61.2026 15.5952L55.0247 1.21483H47.8549V1.46466L47.8871 1.47175C49.1015 1.74322 49.7181 2.14807 49.7181 3.60814V20.7013C49.7181 22.1613 49.0991 22.5662 47.8847 22.8376L47.8525 22.8447V23.0954H52.717V22.8455L52.6848 22.8384C51.4704 22.567 50.8539 22.1621 50.8539 20.702V4.59961L58.79 23.0954H59.2401L67.4074 4.08381V21.1242C67.3033 22.2774 66.6924 22.6334 65.5927 22.8794L65.5602 22.8868V23.1347H74.0338V22.8868L74.0012 22.8794C72.9004 22.6334 72.2747 22.2774 72.1706 21.1242L72.1651 3.60814H72.1706C72.1706 2.14807 72.7872 1.74322 74.0012 1.47175ZM77.8717 12.795C78.0104 9.72381 79.1239 7.50716 80.9926 7.46899C81.5691 7.47843 82.0525 7.66571 82.4283 8.02611C83.2268 8.79371 83.6023 10.3978 83.5443 12.795H77.8717ZM77.7879 13.6606H87.7195V13.6193C87.6913 11.2708 87.0044 9.44407 85.68 8.18978C84.535 7.10584 82.8399 6.50939 81.0594 6.50939H81.0197C80.0956 6.50939 78.9622 6.73129 78.1558 7.13339C77.2377 7.55437 76.4281 8.18388 75.7543 9.01011C74.6698 10.3411 74.0127 12.1399 73.853 14.1615C73.8479 14.222 73.8435 14.2826 73.8391 14.3432C73.8347 14.4038 73.832 14.4581 73.8292 14.5159C73.8236 14.6241 73.8193 14.7327 73.8169 14.8417C73.8129 15.0164 73.8117 15.1923 73.8149 15.3693C73.9225 19.9478 76.4197 23.6068 80.8552 23.6068C84.7484 23.6068 87.0156 20.7882 87.5809 17.0049L87.2952 16.9053C86.3021 18.9386 84.5183 20.1709 82.4883 20.0206C79.717 19.8152 77.5941 17.0324 77.7867 13.6614M98.9485 19.8652C98.6228 20.6304 97.9431 21.0514 97.0321 21.0514C96.1212 21.0514 95.2886 20.4321 94.697 19.3069C94.0614 18.099 93.7269 16.3915 93.7269 14.3688C93.7269 10.159 95.049 7.43594 97.0953 7.43594C97.9522 7.43594 98.6268 7.85693 98.9485 8.59148V19.8652ZM105.537 22.8664C104.322 22.5819 103.706 22.1582 103.706 20.6237V0.208008L96.3274 2.36093V2.62453L96.3727 2.62099C97.3905 2.53955 98.0809 2.67883 98.4818 3.04552C98.7956 3.33273 98.9485 3.77339 98.9485 4.39345V7.1924C98.2212 6.73247 97.3559 6.5086 96.3059 6.5086C94.1766 6.5086 92.2304 7.3966 90.8269 9.00932C89.3638 10.6901 88.5903 12.987 88.5903 15.651C88.5899 20.4089 90.9548 23.6068 94.4746 23.6068C96.5336 23.6068 98.1902 22.4894 98.9485 20.6049V23.1347H105.569V22.8727L105.537 22.8664ZM111.865 2.98453C111.865 1.49889 110.734 0.378368 109.233 0.378368C107.74 0.378368 106.572 1.52329 106.572 2.98453C106.572 4.44578 107.741 5.5907 109.233 5.5907C110.734 5.5907 111.865 4.47017 111.865 2.98453ZM113.606 22.8664C112.392 22.5819 111.775 22.1582 111.775 20.6237H111.77V6.55699L105.149 8.43961V8.69535L105.189 8.69889C106.622 8.82519 107.014 9.31384 107.014 10.9718V23.1347H113.64V22.8727L113.606 22.8664ZM130.572 22.8664C129.357 22.5819 128.741 22.1582 128.741 20.6237V6.55699L122.437 8.37745V8.63398L122.474 8.63791C123.645 8.75988 123.984 9.27647 123.984 10.9419V19.8337C123.593 20.599 122.861 21.0534 121.973 21.0841C120.534 21.0841 119.741 20.1213 119.741 18.3733V6.55739L113.121 8.44001V8.69535L113.16 8.69889C114.593 8.82479 114.985 9.31345 114.985 10.9718V18.4972C114.982 19.0225 115.028 19.547 115.123 20.0639L115.242 20.5773C115.803 22.5666 117.272 23.6068 119.572 23.6068C121.521 23.6068 123.23 22.4119 123.982 20.5423V23.139H130.603V22.877L130.572 22.8664ZM156.292 23.1347V22.8723L156.259 22.8648C154.942 22.5638 154.428 21.9965 154.428 20.8413V11.2728C154.428 8.28932 152.736 6.5086 149.902 6.5086C147.837 6.5086 146.095 7.69089 145.424 9.5326C144.891 7.58112 143.358 6.5086 141.091 6.5086C139.1 6.5086 137.54 7.54925 136.871 9.30716V6.55817L130.25 8.36368V8.62099L130.289 8.62453C131.705 8.74886 132.114 9.25168 132.114 10.8672V23.1347H138.292V22.8727L138.259 22.8648C137.208 22.6201 136.869 22.1739 136.869 21.029V10.0614C137.147 9.41811 137.707 8.65601 138.815 8.65601C140.191 8.65601 140.889 9.60027 140.889 11.4605V23.1347H147.068V22.8727L147.035 22.8648C145.984 22.6201 145.645 22.1739 145.645 21.029V11.2716C145.648 10.9068 145.618 10.5424 145.557 10.1826C145.852 9.48342 146.445 8.65601 147.598 8.65601C148.994 8.65601 149.672 9.57312 149.672 11.4605V23.1347H156.292Z'
        fill={props.fillcolor}
      />
    </g>
    <defs>
      <clipPath id='clip0_406_73'>
        <rect width={156.522} height={24} fill='white' />
      </clipPath>
    </defs>
  </svg>
);
export default MediumSVG;