import { createContextId, Signal } from "@builder.io/qwik";
interface AppState {
  isMenuVisible: boolean;
}
export const AppContext = createContextId<AppState>("appContext");
