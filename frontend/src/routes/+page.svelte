<script>
  import { postRegister, postLogin } from "../lib/apiUser";
  import { success, failed, succesLogin } from "../lib/swall";
  import { goto } from "$app/navigation";
  let statusMenu = false;
  let personLogin = {
    useroremail: "",
    password: "",
  };
  let person = {
    pengguna: "",
    sandi: "",
    email: "",
  };
  function changeMenu() {
    statusMenu = !statusMenu;
  }
  async function submitRegister(e) {
    e.preventDefault();
    const response = await postRegister(person);
    const responseBody = await response.json();
    console.log(responseBody);
    if (responseBody.code === 200) {
      success("Register");
      console.log("done", responseBody.data);
      person = {
        pengguna: "",
        sandi: "",
        email: "",
      };
      await changeMenu();
    } else {
      //   console.log(responseBody.data.join("\n"));
      failed(responseBody.data);
      console.error("gagal datanya", responseBody.data);
    }
  }

  async function submitLogin(e) {
    e.preventDefault();
    const response = await postLogin(personLogin);
    const responseBody = await response.json();
    console.log(responseBody);
    if (response.status === 202 || response.status === 200) {
      succesLogin();
      // console.log("done", responseBody.Token);
      // const dataToke = await responseBody.Token;
      // cookieStore.set("token", dataToke);
      goto("/dashboard");
    } else {
      //   console.log(responseBody.data.join("\n"));
      failed(responseBody.data);
      console.error("gagal datanya", responseBody.data);
    }
  }
</script>

<div class="bg-gray-900 min-h-screen flex justify-center items-center">
  {#if statusMenu === false}
    <div class="text-white border-2 border-amber-500 p-8 rounded-md">
      <h1 class="text-2xl font-bold text-center">Register</h1>
      <form class="py-5 max-w-md" onsubmit={submitRegister}>
        <input
          bind:value={person.pengguna}
          type="text"
          placeholder="Username"
          class="w-full p-2 rounded bg-gray-700 border border-gray-600 text-white mb-4 focus:outline-none focus:ring-2 focus:ring-amber-500"
        />
        <input
          bind:value={person.sandi}
          type="password"
          placeholder="Password"
          class="w-full p-2 rounded bg-gray-700 border border-gray-600 text-white mb-4 focus:outline-none focus:ring-2 focus:ring-amber-500"
        />
        <input
          bind:value={person.email}
          type="email"
          placeholder="Email"
          class="w-full p-2 rounded bg-gray-700 border border-gray-600 text-white mb-4 focus:outline-none focus:ring-2 focus:ring-amber-500"
        />
        <button
          type="submit"
          class="w-full bg-amber-500 hover:bg-amber-600 text-black font-semibold py-2 px-4 rounded"
          >Register</button
        >
      </form>
      <div>
        Jika Sudah Punya Akun dapat
        <button onclick={changeMenu} class="text-amber-400 hover:underline"
          >Login Disini
        </button>
      </div>
    </div>
  {:else}
    <div class="text-white border-2 border-amber-500 p-8 rounded-md">
      <h1 class="text-2xl font-bold text-center">Login</h1>
      <form class="py-5 max-w-md" onsubmit={submitLogin}>
        <input
          bind:value={personLogin.useroremail}
          type="text"
          placeholder="Username/Email"
          class="w-full p-2 rounded bg-gray-700 border border-gray-600 text-white mb-4 focus:outline-none focus:ring-2 focus:ring-amber-500"
        />
        <input
          bind:value={personLogin.password}
          type="password"
          placeholder="Password"
          class="w-full p-2 rounded bg-gray-700 border border-gray-600 text-white mb-4 focus:outline-none focus:ring-2 focus:ring-amber-500"
        />
        <button
          type="submit"
          class="w-full bg-amber-500 hover:bg-amber-600 text-black font-semibold py-2 px-4 rounded"
          >Login</button
        >
      </form>
      <div>
        Belum Punya Akun
        <button onclick={changeMenu} class="text-amber-400 hover:underline"
          >Register Disini
        </button>
      </div>
    </div>
  {/if}
</div>
<!-- <script>
  import "../app.css";
  let username = "";
  let email = "";
  let password = "";

  function register() {
    alert(`Register berhasil!\nUsername: ${username}\nEmail: ${email}`);
    // Bisa tambahkan logic fetch() ke API di sini
  }
</>

<div class="bg-gray-900 min-h-screen flex justify-center items-center">
  <form
    onsubmit={register}
    class="bg-gray-800 text-white p-8 rounded-lg shadow-md w-full max-w-md border border-amber-500"
  >
    <h1 class="text-2xl font-bold mb-6 text-center">Register</h1>

    <div class="block mb-2 text-sm">Username</div>
    <input
      bind:value={username}
      type="text"
      required
      class="w-full p-2 rounded bg-gray-700 text-white mb-4 border border-gray-600 focus:outline-none focus:ring-2 focus:ring-amber-500"
    />

    <div class="block mb-2 text-sm">Email</div>
    <input
      bind:value={email}
      type="email"
      required
      class="w-full p-2 rounded bg-gray-700 text-white mb-4 border border-gray-600 focus:outline-none focus:ring-2 focus:ring-amber-500"
    />

    <div class="block mb-2 text-sm">Password</div>
    <input
      bind:value={password}
      type="password"
      required
      class="w-full p-2 rounded bg-gray-700 text-white mb-6 border border-gray-600 focus:outline-none focus:ring-2 focus:ring-amber-500"
    />

    <button
      type="submit"
      class="w-full bg-amber-500 hover:bg-amber-600 text-black font-semibold py-2 px-4 rounded"
    >
      Register
    </button>
  </form>
</div> -->
