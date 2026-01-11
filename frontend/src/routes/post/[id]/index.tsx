import { component$ } from "@builder.io/qwik";
import { routeLoader$, type DocumentHead } from "@builder.io/qwik-city";
import { Post } from "../post";

const useRouteData = routeLoader$(async (req) => {
  const id = req.params["id"]
  const fetchReq = await fetch(import.meta.env.PUBLIC_SERVER_URL + "/post/" + id)
  const res = await fetchReq.json()
  return res as Post
})
export default component$(() => {
  const post = useRouteData()

  return (

    <main class="flex flex-col">

      <h2 class="text-white text-2xl font-bold text-center">
        {post.value.post.title}
      </h2>

      <div class="flex justify-around align-middle text-white text-sm">

        <p>Posted by {post.value.post.username}</p>

      </div>


      <img

        src={post.value.images[0].imageUrl}
        class="w-full aspect-[4/3] object-contain bg-black mx-auto"
      />


      <div class="mt-4 flex gap-2 overflow-x-auto">
        {post.value.images.map((img, i) =>

          <img
            src={post.value.images[i].imageUrl}
            alt="test"
            class="w-24 thumbnail h-16 object-cover rounded-sm cursor-pointer border-2 border-transparent hover:border-blue-500"
          />
        )}
      </div>
      {/* 
      
      <div class="lg:w-5xl w-full">
      
      
        <form
          hx-post="/submit-comment/{{.Post.ID}}"
          hx-swap="beforeend"
          hx-target="#comments"
          class="my-4"
        >
          <textarea
            name="content"
            required
      
            placeholder="Write a comment"
            class="w-full p-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
          ></textarea>
          <button
            type="submit"
            class="mt-2 px-3 py-2 bg-slate-500 text-white rounded-sm hover:bg-cyan-300 hover:text-slate-500 cursor-pointer"
          >
            Post Comment
          </button>
        </form>
      

     
        <div
          id="comments"
          hx-get="/get-post-comments/{{.Post.ID}}"
          hx-trigger="load"
          hx-swap="innerHTML"
          class="space-y-4 text-slate-50"
        >
     
        </div>
      </div> */}

    </main>

  );
});

export const head: DocumentHead = {
  title: "Nub Global",
  meta: [
    {
      name: "description",
      content: "Curious?",
    },
  ],
};

