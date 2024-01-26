import { NextResponse } from 'next/server'

// TODO: implement this if everything else works out
// This function can be marked `async` if using `await` inside
export async function middleware(request) {

    try {
        const response = await fetch(process.env.NEXT_PUBLIC_ENDPOINT + "/auth", {
            method: "POST",
            credentials: "include"
        })

        console.log(response)
        if (response.ok) {
            return NextResponse.redirect(request.url)
        } else {
            return NextResponse.redirect(new URL('/notfound', request.url))
        }
    } catch (error) {
        return NextResponse.redirect(new URL('/notfound', request.url))
    }

    // return NextResponse.redirect(new URL('/home', request.url))
}

// See "Matching Paths" below to learn more
export const config = {
    matcher: '/awdasdsad',
}