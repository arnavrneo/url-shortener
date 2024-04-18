"use client";

import React, {useEffect, useState} from 'react'
import FourOOne from "@/pages/Unauthorized";
import {useRouter} from "next/navigation";

function Main() {
    const [userName, setUserName] = useState('');
    const [logged, setLogged] = useState(false);
    const [url, setUrl] = useState('');
    const [visible, setVisible] = useState(false);
    const [shortLink, setShortLink] = useState('');

    const router = useRouter();

    useEffect(() => {
        (
            async () => {
                const response = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL + '/user', {
                    credentials: 'include',
                    headers: {"Content-Type": "application/json"},
                });

                if (response.ok) {
                    const content = await response.json();
                    setUserName(content.username);
                    setLogged(true)
                } else {
                    setLogged(false)
                }
            }
        )();
    }, []);


    const handleShorten = async (e) => {
        e.preventDefault();

        const formData = new FormData();
        formData.append("url", url);

        try {
            const response = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL + "/shorten", {
                method: "POST",
                credentials: "include",
                // headers: {"Content-Type": "application/json"},
                body: formData,
            })

            if (response.ok) {
                const data = await response.json()
                console.log(data)
                setShortLink(data.shorten_link)
                setVisible(true)
            } else {
                alert("url cannot be sent")
            }
        } catch (error) {
            console.log(error);
        }
    }

    const handleLogout = async () => {
        try {
            const response = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL + "/logout", {
                credentials: "include",
                method: "POST",
            })

            if (response.ok) {
                console.log("logged out successfully")
                router.push("/")
            }
        } catch (error) {
            console.log(error)
        }
    }

    return (
        <div>
            {logged ? <div className="flex flex-wrap min-h-screen w-full content-center justify-center bg-gray-200 py-10">
                <div className="flex shadow-md">
                    <div className="flex flex-wrap content-center justify-center rounded-l-md bg-white" style={{ width: "24rem", height: "32rem" }}>
                        <div className="w-72">
                            <div className="flex flex-wrap content-center justify-center ">
                                <h1 className="text-xl font-semibold">url-shortener</h1>
                            </div>
                            <small className="text-gray-400 mt-4 flex flex-wrap content-center justify-center ">Welcome {userName}</small>
                            <small className="text-gray-400 mt-4 flex flex-wrap content-center justify-center ">Short your url!</small>


                            <form onSubmit={handleShorten} className="mt-4">
                                <div className="mb-3">
                                    <label className="mb-2 block text-xs font-semibold">Your url  </label>
                                    <input onChange={(e) => setUrl(e.target.value)} type="url" name="url" placeholder="Enter a URL" className="block w-full rounded-md border border-gray-300 focus:border-purple-700 focus:outline-none focus:ring-1 focus:ring-purple-700 py-1 px-1.5 text-gray-500" required/>
                                </div>
                                <div className="mb-3">
                                    <button className="mb-1.5 block w-full text-center text-white bg-purple-700 hover:bg-purple-900 px-2 py-1.5 rounded-md">Make it short!</button>
                                </div>
                            </form>

                            <div className="text-center">
                                <span className="text-xs text-gray-400 font-semibold"> by </span>
                                <a href="https://www.github.com/arnavrneo" className="text-xs font-semibold text-purple-700">arnavrneo</a>
                            </div>
                            <div className="text-center mt-12">
                                <div className="flex flex-row justify-center mb-3">
                                    <button onClick={handleLogout} className="mb-1.5 block w-min text-center text-white bg-purple-700 hover:bg-purple-900 px-2 py-1.5 rounded-md">Logout</button>
                                </div>
                            </div>

                            {/*LINK VISIBLE HERE*/}
                            {visible ?
                                <div className="text-center">
                                    <span className="text-xs text-gray-400 font-semibold"> by </span>
                                    <a href={shortLink} className="text-xs font-semibold text-purple-700">{shortLink}</a>
                                </div> : ''}
                        </div>
                    </div>


                </div>

                <div className="mt-3 w-full">
                    <p className="text-center">frontend for <span className="text-purple-700">url-shortener </span>(in nextjs)</p>
                </div>
            </div> : <FourOOne />}
        </div>
    )
}

export default Main;