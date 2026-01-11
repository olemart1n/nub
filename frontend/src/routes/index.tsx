import { $, component$, useOnDocument, useStore } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import { Post1, type Post1Type } from "~/components/post-1";

export default component$(() => {

  const store = useStore<Post1Type[]>([])

  useOnDocument("DOMContentLoaded", $(async () => {

    const req = await fetch(import.meta.env.PUBLIC_SERVER_URL + "/latest-images/0")
    const res = await req.json()
    console.log(res)
    res.forEach((post: Post1Type) => (

      store.push({
        id: post.id.toString(),
        postId: post.postId.toString(),
        country: post.country,
        createdAt: post.createdAt,
        imageUrl: post.imageUrl,
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
          return <Post1 id={image.id} postId={image.postId} createdAt={image.createdAt} country={image.country} imageUrl={image.imageUrl} />
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

