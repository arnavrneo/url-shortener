export default function FourOOne() {
    return  <>
        <div className="flex flex-col items-center justify-center h-screen bg-gradient-to-r from-sky-500 to-indigo-500">
            <h1 className="text-5xl font-bold text-white mb-8">Restricted Access</h1>
            <p className="text-lg text-gray-200 mb-8">
                You are unauthenticated. Login or return to the previous page.
            </p>
            <div className="flex flex-wrap justify-center gap-4">
                <a
                    href="/"
                    className="inline-flex items-center px-4 py-2 text-base font-medium text-white bg-indigo-600 border border-transparent rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                    Go Back
                </a>
            </div>
        </div>
    </>
}