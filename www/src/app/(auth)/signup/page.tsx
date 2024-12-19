'use client';

import Link from 'next/link';

import { z } from 'zod';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';

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
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { toast } from 'sonner';
import { fetchWithAuth } from '@/lib/fetcher';

const formSchema = z.object({
  name: z.string().min(3, {
    message: 'Name must be at least 3 character',
  }),
  email: z.string().email({
    message: 'Invalid email address',
  }),
  password: z.string().min(8, {
    message: 'Password must be at least 8 character',
  }),
});

export default function SignUp() {
  const form = useForm({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: '',
      email: '',
      password: '',
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      const res = await fetchWithAuth('/users/create', {
        method: 'POST',
        body: JSON.stringify(values),
      });

      if (!res.ok) {
        toast.error('Account creation failed');
        throw new Error(`Error: ${res.status} - ${res.statusText}`);
      }

      toast.success('Account created successfully');

      const data = await res.json();
      console.log('Data: ', data);
    } catch (error) {
      console.error('Error submitting the form:', error);
    }
  }

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
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <CardContent className='space-y-2'>
              <FormField
                control={form.control}
                name='name'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Name</FormLabel>
                    <FormControl>
                      <Input
                        placeholder='Enter your Name'
                        className='rounded focus-visible:ring-0'
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='email'
                render={({ field }) => (
                  <FormItem>
                    <Label htmlFor='email'>E-Mail</Label>
                    <FormControl>
                      <Input
                        placeholder='Enter your E-Mail'
                        className='rounded focus-visible:ring-0'
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='password'
                render={({ field }) => (
                  <FormItem>
                    <Label htmlFor='password'>Password</Label>
                    <FormControl>
                      <Input
                        type='password'
                        placeholder='Enter your Password'
                        className='rounded focus-visible:ring-0'
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <Button type='submit' variant='amber' className='!mt-4 w-full'>
                Log In
              </Button>
            </CardContent>
          </form>
        </Form>
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
