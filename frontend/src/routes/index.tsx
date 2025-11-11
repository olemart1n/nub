import { $, component$, useOnDocument, useSignal, useStore } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import { Post1, type Post1Type } from "~/components/post-1";

export default component$(() => {

  const store = useStore<Post1Type[]>([])


  useOnDocument("DOMContentLoaded", $(async () => {

    const req = await fetch(import.meta.env.PUBLIC_SERVER_URL + "/get-latest-images/0")
    const res = await req.json()
    res.forEach((post: any) => (

      store.push({
        id: post.ID.toString(),
        postId: post.PostID.toString(),
        country: post.Country,
        createdAt: post.CreatedAt,
        imageURL: post.ImageURL,
      })
    ));
  }))

  return (
    <main class="p-2">
      <h1 class="my-3 sr-only">Below is the latest uploads..</h1>

      {/* {{template "filter-options" .}} */}
      <div
        class="grid grid-cols-2 md:grid-cols-4 gap-4 w-full"
      >
        {store.map((image) => {
          return <Post1 id={image.id} postId={image.postId} createdAt={image.createdAt} country={image.country} imageURL={image.imageURL} />
        })}

      </div>
      <div class="h-20 w-full p-2 text-center text-black">
        <button
          type="button"
          class="bg-yellow-300 p-1 text-sm cursor-pointer"
          id="load-more-posts-btn"
        >
          See more posts
        </button>
      </div>
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

