import { component$, Slot, useContextProvider, useStore } from '@builder.io/qwik';
import { Header } from '~/components/header';
import { AppContext } from '~/context';
export default component$(() => {

  const appState = useStore({
    isMenuVisible: false
  })
  useContextProvider(AppContext, appState)
  return (
    <>
      <Header />
      <Slot /> {/* <== This is where the route will be inserted */}
      <footer>
        <p>© Nub Global — Preserving the Past</p>
      </footer>
    </>
  );
});