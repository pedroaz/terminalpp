<script lang="ts">
  import AppSidebar from "$lib/components/app-sidebar.svelte";
  // import { Button } from "$lib/components/ui/button";
  import * as Sidebar from "$lib/components/ui/sidebar/index.js";
  import Counter from "./lib/Counter.svelte";
  import "./app.css";
</script>

<Sidebar.Provider>
  <AppSidebar />
  <main>
    <Sidebar.Trigger />
  </main>
</Sidebar.Provider>

<!-- <main>
  <AppSidebar />
</main> -->
