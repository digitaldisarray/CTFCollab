<script lang="ts">
    import DataTable from "./data-table.svelte";
    import { columns } from "./columns.js";
    import { data } from "./columns.js";
  </script>
  
  <div class="page-container">
    <div class="container">
      <!-- Header -->
      <header>
        <h1>Event Challenges</h1>
        <p>Navigate Event Challenges here.</p>
      </header>
  
      <!-- DataTable -->
      <DataTable {data} {columns} />
    </div>
    <div class="container">
        <!-- Header -->
        <header>
          <h1>Team Members</h1>
          <p>Invite team members to collaborate on challenges.</p>
        </header>
    
      </div>
  </div>
  
  <style>
    .page-container {
      display: flex;
      justify-content: flex-start; /* Align container to the left */
      width: 100%; /* Ensure the page container takes full width */
    }
  
    .container {
      display: flex;
      flex-direction: column;
      width: 50%; /* Set container width to 50% of the screen */
      padding: 20px;
      box-sizing: border-box; /* Include padding in the width calculation */
    }
  
    header h1 {
      font-size: 2.5rem;
      color: #333;
      margin-bottom: 10px;
    }
  
    header p {
      font-size: 1.2rem;
      color: #666;
    }
  </style>