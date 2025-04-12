import { NextRequest, NextResponse } from 'next/server';
import { getSessionCookie } from 'better-auth/cookies';

export async function middleware(request: NextRequest) {
  const cookies = getSessionCookie(request);
  return NextResponse.next();
}
