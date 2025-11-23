<script>
  import { page } from "$app/stores";
  import { postLogout } from "../../lib/apiUser";
  import { goto } from "$app/navigation";
  // reactive store untuk ambil path aktif
  let pathname = $derived($page.url.pathname);
  let { children } = $props();
  let sidebarOpen = $state(true);
  let showProfileMenu = $state(false);

  let user = {
    name: "Admin POS",
    email: "admin@pos.com",
    avatar: "https://ui-avatars.com/api/?name=Admin+POS",
  };

  let menus = [
    { name: "Dashboard", href: "/dashboard", icon: "📊" },
    { name: "Produk", href: "/product", icon: "📦" },
    { name: "Transaksi", href: "/transaksi", icon: "🛒" },
    { name: "Laporan", href: "/laporan", icon: "📑" },
    { name: "Pengaturan", href: "/setting", icon: "⚙️" },
  ];

  function toggleSidebar() {
    sidebarOpen = !sidebarOpen;
  }

  async function logout() {
    const response = await postLogout();
    const responseBody = await response.json();
    console.log(responseBody);
    if (response.status === 202 || response.status === 200) {
      console.log(responseBody);
      goto("/");
    } else {
      failed(responseBody);
      console.log("gagal datanya", responseBody);
    }
  }
</script>

<div class="flex h-screen bg-gray-100">
  <!-- Sidebar -->
  <aside
    class={`${
      sidebarOpen ? "w-64" : "w-16"
    } bg-white shadow-lg transition-all duration-300 flex flex-col`}
  >
    <div class="flex items-center justify-between p-4 border-b">
      <span class={`font-bold text-lg ${!sidebarOpen ? "hidden" : ""}`}
        >POS System</span
      >
      <button class="p-2 rounded hover:bg-gray-200" onclick={toggleSidebar}>
        {sidebarOpen ? "⬅️" : "➡️"}
      </button>
    </div>

    <nav class="mt-4 flex-1">
      {#each menus as menu}
        <a
          href={menu.href}
          class="flex items-center gap-2 px-4 py-2 text-gray-700 hover:bg-gray-200 rounded-md
                 {pathname === menu.href ? 'bg-gray-200 font-semibold' : ''}"
        >
          <span>{menu.icon}</span>
          {#if sidebarOpen}
            <span>{menu.name}</span>
          {/if}
        </a>
      {/each}
    </nav>
  </aside>

  <!-- Main -->
  <main class="flex-1 flex flex-col">
    <!-- Navbar -->
    <header class="h-14 bg-white shadow flex items-center justify-between px-6">
      <h1 class="text-lg font-semibold">Dashboard</h1>

      <!-- Profile -->
      <div class="relative">
        <button
          class="flex items-center gap-2"
          onclick={() => (showProfileMenu = !showProfileMenu)}
        >
          <img src={user.avatar} alt="avatar" class="w-8 h-8 rounded-full" />
          <span class="hidden md:inline">{user.name}</span>
        </button>

        {#if showProfileMenu}
          <div
            class="absolute right-0 mt-2 w-48 bg-white border rounded-md shadow-lg z-10"
          >
            <div class="p-2 border-b">
              <p class="font-semibold">{user.name}</p>
              <p class="text-sm text-gray-500">{user.email}</p>
            </div>
            <button
              class="block w-full text-left px-4 py-2 hover:bg-gray-100"
              onclick={logout}
            >
              Logout
            </button>
          </div>
        {/if}
      </div>
    </header>

    <!-- Content -->
    <section class="p-6 overflow-y-auto">
      {@render children()}
    </section>
  </main>
</div>
