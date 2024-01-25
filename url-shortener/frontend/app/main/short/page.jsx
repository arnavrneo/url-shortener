import React, {useEffect, useState} from 'react'
import process from "next/dist/build/webpack/loaders/resolve-url-loader/lib/postcss";

function Short() {
  const [userData, setUserData] = useState('');
  const [logged, setLogged] = useState(false);

  useEffect(() => {
    (
        async () => {
          const response = await fetch(process.env.NEXT_PUBLIC_ENDPOINT + '/user', {
            credentials: 'include',
            headers: {"Content-Type": "application/json"},
          });

          if (response.ok) {
            const content = await response.json();
            setUserData(content.message);
            setLogged(true)
          } else {
            setLogged(false)
          }
        }
    )();
  }, []);

  return (
    <div>
      {logged ? <div className="flex flex-wrap min-h-screen w-full content-center justify-center bg-gray-200 py-10">

        <div className="flex shadow-md">

          <div className="flex flex-wrap content-center justify-center rounded-l-md bg-white" style={{ width: "24rem", height: "32rem" }}>
            <div className="w-72">
              <div className="flex flex-wrap content-center justify-center ">
                <h1 className="text-xl font-semibold">url-shortener</h1>
              </div>
              <small className="text-gray-400 m-10">Enter your url and make it short!</small>

              <div className="m-3 flex flex-wrap content-center justify-center">
                <h3 className="mb-2 block font-semibold">Your short link  </h3>
              </div>
              <div className="mb-3 flex flex-wrap content-center justify-center">
                <h3 className="mb-2 block font-semibold">
                  https://short-link.com
                </h3>
              </div>

              <div className="text-center">
                <span className="text-xs text-gray-400 font-semibold"> by </span>
                <a href="https://www.github.com/arnavrneo" className="text-xs font-semibold text-purple-700">arnavrneo</a>
              </div>
            </div>
          </div>
        </div>
        <div className="mt-3 w-full">
          <p className="text-center">frontend for <span className="text-purple-700">url-shortener </span>(in nextjs)</p>
        </div>
      </div> : ''}
    </div>
  )
}

export default Short;