<script>
  import { getProduct } from "../../../lib/apiProduct";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { postLogout } from "../../../lib/apiUser";
  import { fetchApi } from "$lib/apiAuth";

  let datas = $state([]);
  async function getProducts() {
    const response = await getProduct();
    const responseBody = await response.json();
    fetchApi(response.status);
    datas = responseBody.data; // data array dari backend
    console.log(responseBody);
  }
  onMount(async () => {
    await getProducts();
  });
</script>

<div class="p-6">
  <h1 class="text-2xl font-bold mb-4">Daftar Produk</h1>
  <div class="overflow-x-auto">
    <button
      class="bg-blue-600 rounded-md p-2 text-amber-50 hover:bg-blue-800 transition duration-700"
      >Tambah Product</button
    >
    <table
      class="min-w-full bg-white border border-gray-200 rounded-lg shadow-md"
    >
      <thead class="bg-gray-100">
        <tr>
          <th class="px-4 py-2 text-left border-b">No</th>
          <th class="px-4 py-2 text-left border-b">Kode Barang</th>
          <th class="px-4 py-2 text-left border-b">Nama Barang</th>
          <th class="px-4 py-2 text-left border-b">Harga Beli</th>
          <th class="px-4 py-2 text-left border-b">Harga Jual</th>
          <th class="px-4 py-2 text-left border-b">Profit</th>
          <th class="px-4 py-2 text-left border-b">Keterangan</th>
          <th class="px-4 py-2 text-left border-b">Stok</th>
          <th class="px-4 py-2 text-left border-b">Action</th>
        </tr>
      </thead>
      <tbody>
        {#each datas as product, index (product.id)}
          <tr class="hover:bg-gray-50">
            <td class="px-4 py-2 border-b">{index + 1}</td>
            <td class="px-4 py-2 border-b">{product.kodebarang}</td>
            <td class="px-4 py-2 border-b">{product.nameprod}</td>
            <td class="px-4 py-2 border-b">Rp {product.HargaProd}</td>
            <td class="px-4 py-2 border-b">Rp {product.jualprod}</td>
            <td class="px-4 py-2 border-b">Rp {product.profitprod}</td>
            <td class="px-4 py-2 border-b">{product.keterangan}</td>
            <td class="px-4 py-2 border-b">{product.stok}</td>
            <td class="px-4 py-2 border-b">
              <button
                class="bg-green-700 border rounded-md w-16"
                onclick={() => handleEdit(item.id, item.name, item.status)}
                >Edit</button
              >
              <button
                class="bg-red-700 border rounded-md w-16"
                onclick={() => handleDelete(item.id)}>Delete</button
              >
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>
