'use client';

import { BorderBeam } from '@/components/magicui/border-beam';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { authClient } from '@/lib/auth-client';
import { zodResolver } from '@hookform/resolvers/zod';
import { ReloadIcon } from '@radix-ui/react-icons';
import Link from 'next/link';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import { z } from 'zod';

const formSchema = z.object({
  email: z.string().email({
    message: 'Invalid email address',
  }),
  password: z.string().min(8, {
    message: 'Password must be at least 8 character',
  }),
});

export default function SignIn() {
  const [loading, setLoading] = useState(false);
  const form = useForm({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      const { data, error } = await authClient.signIn.email(
        {
          email: values.email,
          password: values.password,
          callbackURL: '/',
        },
        {
          onRequest: (ctx) => {
            setLoading(true);
          },
          onSuccess: (ctx) => {
            setLoading(false);
            toast.success('Logged in successfully');
            console.log(ctx.response);
          },
          onError: (ctx) => {
            setLoading(false);
            toast.error(ctx.error.message);
            console.log(ctx.error);
          },
        }
      );
    } catch (error) {
      console.error('Error submitting the form:', error);
    }
  }

  return (
    <div className='mx-4 w-full md:mx-auto md:w-1/2 lg:mx-auto lg:w-1/2'>
      <Card className='relative overflow-hidden rounded-md shadow-none'>
        <CardHeader>
          <CardTitle>Sign In</CardTitle>
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
              <Button
                type='submit'
                variant='amber'
                className='mt-4! w-full'
                disabled={loading}
              >
                {loading ? 'Logging...' : 'Log In'}
              </Button>
            </CardContent>
          </form>
        </Form>
        <CardFooter className='text-sm'>
          Don't have an account ?
          <Link href='/signup' className='ml-1 underline'>
            Sign Up
          </Link>
        </CardFooter>
        <BorderBeam duration={8} size={100} colorTo='#ffd230' />
      </Card>
    </div>
  );
}
