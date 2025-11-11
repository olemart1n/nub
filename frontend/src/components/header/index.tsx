import { component$, useContext, useSignal, useVisibleTask$ } from "@builder.io/qwik";
import { AppContext } from "~/context";


export const Header = component$(() => {
  const app = useContext(AppContext)
  const isLoggedIn = false

  return (
    <header class="w-full p-1">
      <nav class="flex w-full relative justify-between">
        <a
          href="/"
          class="text-black dark:text-yellow-500 hover:bg-yellow-100 block my-auto"
        >www.nub.global
        </a>

        <button
          onClick$={() => {
            app.isMenuVisible = !app.isMenuVisible
          }}
          class="flex flex-col justify-center z-10 items-center min-w-8 min-h-10 space-y-1 cursor-pointer"
        >
          <span
            class="block w-6 h-0.5 bg-slate-950 dark:bg-slate-200 transition-transform duration-300 group-hover:translate-x-1"
          ></span>
          <span
            class="block w-6 h-0.5 bg-slate-950 dark:bg-slate-200 transition-transform duration-300 group-hover:translate-x-2"
          ></span>
          <span
            class="block w-6 h-0.5 bg-slate-950 dark:bg-slate-200 transition-transform duration-300 group-hover:translate-x-3"
          ></span>
        </button>

        <div
          class="fixed sm:-right-1/4 -right-1/2 h-full w-full sm:w-1/2 text-white"
        >

          <div
            id="menu"
            class={"relative flex h-full w-1/2 gap-y-3 py-10 transform flex-col bg-slate-800/[.6] duration-300 ease-in-out " + (app.isMenuVisible ? "" : " z-50 translate-x-full")}
          >


            {isLoggedIn ?
              <>
                <a
                  class="ml-auto px-3 mx-2 bg-slate-600 w-full py-2 text-center hover:bg-slate-500"
                  href="/upload"
                >Upload</a>

                <a
                  class="ml-auto relative px-3 mx-2 bg-slate-600 w-full py-2 text-center hover:bg-slate-500 flex items-center justify-center gap-2"
                  href="/profile"
                >
                  Profile</a>

                <button
                  hx-post="/sign-out"
                  class="ml-auto px-3 mx-2 x-30 bg-slate-600 w-full py-2 text-center hover:bg-slate-500 cursor-pointer"
                  hx-redirect="/"
                >
                  Sign out
                </button>
              </>
              : <a
                class="ml-auto px-3 mx-2 bg-slate-600 w-full py-2 text-center hover:bg-slate-500"
                href="/upload"
              >Upload</a>}



          </div>
        </div>
      </nav>
    </header>


  )
})


{/* <!-- Profile Icon
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            class="w-5 h-5 text-white"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5.121 17.804A9.004 9.004 0 0112 15c2.21 0 4.21.805 5.879 2.121M15 11a3 3 0 11-6 0 3 3 0 016 0z"
            />
          </svg>
          --> */}