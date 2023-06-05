import { Component, For, createSignal, createEffect, Show } from "solid-js";
import {
  GetGamesList,
  SetGame,
  CheckConn,
  Reconnect,
  PinGame,
  GetPins,
  IsMac,
} from "../wailsjs/go/main/App";
import { faToggleOn, faThumbTack } from "@fortawesome/free-solid-svg-icons";
import Fa, { FaLayers } from "solid-fa";

const App: Component = () => {
  const [gamesList, setGamesList] = createSignal([
    { title: "Home", img: "home" },
  ]);
  const [pinsShow, setPinsShow] = createSignal(false);
  const [selection, setSelection] = createSignal("Home");
  const [status, setStatus] = createSignal("Online");
  const [connErr, setConnErr] = createSignal(false);
  const [isMac, setIsMac] = createSignal(false);

  GetGamesList().then((result: string) => setGamesList(gamesList().concat(JSON.parse(result))));
  IsMac().then((result: boolean) => setIsMac(result));

  const connCheck = () => {
    CheckConn().then((result: boolean) => {
      if (result) setConnErr(true);
    });
  };

  createEffect(() => {
    selection();
    status();
    connCheck();
  });

  return (
    <div class="text-white text-center select-none">
      <div class={`bg-red-600 ${isMac() ? "pt-16" : "pt-6"} pb-6`}>
        <p class=" text-2xl font-semibold">[NS-RPC]</p>
        <FaLayers size="2x">
          <Fa icon={faToggleOn} />
        </FaLayers>
      </div>
      <div class="container pt-5 pb-5">
        <label for="games" class="block mb-2 font-medium">
          {!pinsShow() ? "Game" : "Pins"}
        </label>
        <select
          id="games"
          class="bg-slate-800 border border-white rounded-lg focus:border-red-600 w-80 h-10"
          onChange={(e) => setSelection(e.currentTarget.value)}
          onMouseOver={() => {
            if (gamesList().length < 2) {
              GetGamesList().then((result: string) =>
                setGamesList(JSON.parse(result))
              );
            }
          }}
        >
          <For each={gamesList()}>
            {(game: { title: string; img: string }) => (
              <option value={game.title}>{game.title}</option>
            )}
          </For>
        </select>
        <label for="status" class="block pt-5 mb-2 font-medium">
          Status
        </label>
        <input
          id="status"
          class="bg-slate-800 border border-white rounded-lg focus:border-red-600 w-80 h-10 pl-2 pr-2"
          onChange={(e) => setStatus(e.currentTarget.value)}
          placeholder="Online, Karting with Friends, etc..."
        />
      </div>
      <button
        class="rounded-xl bg-red-700 w-20 h-10"
        onClick={() => SetGame(selection(), status())}
      >
        Play
      </button>
      <button
        class="ml-2 rounded-xl bg-yellow-400 text-black w-20 h-10"
        onClick={() => SetGame("Home", "Idle")}
      >
        Idle
      </button>
      <div class="mt-2" />
      <button
        class="rounded-xl bg-indigo-700 w-10 h-10"
        onClick={() => PinGame(selection())}
      >
        <FaLayers>
          <Fa icon={faThumbTack} />
        </FaLayers>
      </button>
      <button
        class="ml-2 rounded-xl bg-indigo-700 w-24 h-10"
        onClick={() => {
          if (!pinsShow())
            GetPins().then((result: string) => {
              setGamesList(JSON.parse(result));
            });
          else
            GetGamesList().then((result: string) =>
              setGamesList(JSON.parse(result))
            );
          setPinsShow(!pinsShow());
        }}
      >
        Switch {!pinsShow() ? "Pins" : "List"}
      </button>
      <Show when={connErr()}>
        <p
          onClick={() =>
            Reconnect().then((result: boolean) => {
              if (result) setConnErr(false);
            })
          }
          class="mt-5 font-mono underline bg-red-600"
        >
          Couldn't hook the Discord client.
          <br />
          Ensure Discord is started, then click this message to retry.
        </p>
      </Show>
    </div>
  );
};

export default App;
