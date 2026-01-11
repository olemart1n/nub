export interface Post {
  images: Image[];
  post: {
    createdAt: string;
    id: number;
    title: string;
    userID: number;
    username: string;
  };
}

export interface Image {
  id: number;
  postID: number;
  country: string;
  imageUrl: string;
  createdAt: string;
}
