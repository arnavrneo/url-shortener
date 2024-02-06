"use client"

import React, {useState} from "react";
import {useRouter} from "next/navigation";

export default function Login() {
  // form field name depends on these names
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('')
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch(process.env.NEXT_PUBLIC_ENDPOINT + '/login', {
        method: "POST",
        credentials: "include",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
              username,
              email,
              password,
            }),
      });
      console.log(JSON.stringify({
        username,
        email,
        password,
      }))
      if (response.ok) { // response.ok => for 200 http code; not data.ok
        await router.push('/main');
      } else {
        alert("Invalid credentials");
      }
    } catch (error) {
      console.log(error)
    }
  };

  return (
    <div>
      <div className="flex flex-wrap min-h-screen w-full content-center justify-center bg-gray-200 py-10">

        <div className="flex shadow-md">

          <div className="flex flex-wrap content-center justify-center rounded-l-md bg-white" style={{width: "24rem", height: "32rem"}}>
            <div className="w-72">

              <h1 className="text-xl font-semibold">Login</h1>
              <small className="text-gray-400">Welcome! Please enter your details</small>

              <form onSubmit={handleSubmit} className="mt-4">
                <div className="mb-3">
                  <label className="mb-2 block text-xs font-semibold">Name</label>
                  <input onChange={(e) => setUsername(e.target.value)} name="username" type="username" placeholder="Your name" className="block w-full rounded-md border border-gray-300 focus:border-purple-700 focus:outline-none focus:ring-1 focus:ring-purple-700 py-1 px-1.5 text-gray-500" required />
                </div>
                <div className="mb-3">
                  <label className="mb-2 block text-xs font-semibold">Email</label>
                  <input onChange={(e) => setEmail(e.target.value)} type="email" placeholder="Enter your email" className="block w-full rounded-md border border-gray-300 focus:border-purple-700 focus:outline-none focus:ring-1 focus:ring-purple-700 py-1 px-1.5 text-gray-500" required/>
                </div>

                <div className="mb-3">
                  <label className="mb-2 block text-xs font-semibold">Password</label>
                  <input onChange={(e) => setPassword(e.target.value)} type="password" placeholder="*******" className="block w-full rounded-md border border-gray-300 focus:border-purple-700 focus:outline-none focus:ring-1 focus:ring-purple-700 py-1 px-1.5 text-gray-500" required/>
                </div>

                {/* FORGOT PASSWORD
                <div className="mb-3 flex flex-wrap content-center">
                  <input id="remember" type="checkbox" className="mr-1 checked:bg-purple-700" /> <label for="remember" className="mr-auto text-xs font-semibold">Remember for 30 days</label>
                  <a href="#" className="text-xs font-semibold text-purple-700">Forgot password?</a>
                </div> */}


                <div className="mb-3">
                  <button type="submit" className="mb-1.5 block w-full text-center text-white bg-purple-700 hover:bg-purple-900 px-2 py-1.5 rounded-md">Sign in</button>
                  {/* GOOGLE IDP
                  <button className="flex flex-wrap justify-center w-full border border-gray-300 hover:border-gray-500 px-2 py-1.5 rounded-md">
                    <img className="w-5 mr-2" src="https://lh3.googleusercontent.com/COxitqgJr1sJnIDe8-jiKhxDx1FrYbtRHKJ9z_hELisAlapwE9LUPh6fcXIfb5vwpbMl4xl9H9TRFPc5NOO8Sb3VSgIBrfRYvW6cUA"/>
                      Sign in with Google
                  </button> */}
                </div>
              </form>

              <div className="text-center">
                <span className="text-xs text-gray-400 font-semibold">Don't have account? </span>
                <a href="/signup" className="text-xs font-semibold text-purple-700">Sign up</a>
              </div>
            </div>
          </div>


        </div>

        <div className="mt-3 w-full">
          <p className="text-center">frontend for <span className="text-purple-700">url-shortener </span>(in nextjs)</p>
        </div>
      </div>
    </div>
  );
}
