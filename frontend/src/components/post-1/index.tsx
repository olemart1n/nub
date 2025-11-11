import { component$ } from "@builder.io/qwik";


export interface Post1Type {
  country: string;
  createdAt: string;
  postId: string;
  imageURL: string;
  id: string
}


export const Post1 = component$<Post1Type>(({ id, postId, country, createdAt, imageURL }) => {
  return (
    <a
      id={id}
      href={"/post/" + postId}
      class={"overflow-hidden relative flex aspect-square flex-col rounded shadow-lg hover:shadow-xl duration-200 w-full aspect-square border border-slate-800 hover:border-slate-600 " + (imageURL.length === 0 ? "wave-animation" : "")}
    >
      <div class="flex justify-between items-center">
        <span class="bg-yellow-700 whitespace-nowrap text-sm text-slate-200 p-1 inline z-10 py-0 outline-0.5"
        >{country}</span>
      </div>
      <img
        src={imageURL}
        alt="Image"
        class="object-cover flex-1 w-full h-3/4"
        loading="lazy"
      />
      <p class="italic text-start text-slate-400 p-1 text-sm">
        Posted: {createdAt}
      </p>
    </a>

  )
})