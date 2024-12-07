import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import Link from 'next/link';

export default function SignUp() {
  return (
    <div className='mx-4 w-full md:mx-auto lg:mx-auto lg:w-1/2'>
      <Card className='rounded-md shadow-none'>
        <CardHeader>
          <CardTitle>Sign Up</CardTitle>
          <CardDescription>
            Create an account and discover the worlds' best UI component
            framework.
          </CardDescription>
        </CardHeader>
        <CardContent className='space-y-2'>
          <div className='space-y-1'>
            <Label htmlFor='name'>Name</Label>
            <Input placeholder='Enter your Name' className='rounded focus-visible:ring-0'  />
          </div>
          <div className='space-y-1'>
            <Label htmlFor='email'>E-Mail</Label>
            <Input placeholder='Enter your E-Mail' className='rounded focus-visible:ring-0' />
          </div>
          <div className='space-y-1'>
            <Label htmlFor='email'>Password</Label>
            <Input placeholder='Enter your Password' className='rounded focus-visible:ring-0'  />
          </div>
          <Button variant='amber' className='!mt-4 w-full'>
            Log In
          </Button>
        </CardContent>
        <CardFooter className='text-sm'>
          Already have an account ?
          <Link href='/signin' className='ml-1 underline'>
            Sign In
          </Link>
        </CardFooter>
      </Card>
    </div>
  );
}
