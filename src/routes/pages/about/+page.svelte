<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input/index.js";
    import { goto } from '$app/navigation';
</script>

<header>
</header>

<main>
    
    <div class="welcome-container">
    <div class="absolute left-4 top-4 md:left-8 md:top-7">
        <a href="/" class="logo">
            <span class="ctf">CTF</span>
            <span class="collab">Collab</span>
        </a>
    </div>

    <div class="absolute right-4 top-4 md:right-8 md:top-8">
        {#if data.user}
            <Button href="/pages/admin-page" variant="ghost">Dashboard</Button>
        {:else}
            <Button href="/pages/signin" variant="ghost">Login</Button>
        {/if}
        <span class="separator">|</span>
        <Button href="/pages/about" variant="ghost">About</Button>
    </div>
    
    <div class="about-container">
        <div class="features">
            <div class="feature-card">
                <div class="feature-title">
                    <div class="feature-icon">🏆</div>
                    Competition Management
                </div>
                <p>Keep track of upcoming, current, and past CTF events your team is participating in. Never miss a competition deadline again.</p>
            </div>
            
            <div class="feature-card">
                <div class="feature-title">
                    <div class="feature-icon">🎯</div>
                    Challenge Tracking
                </div>
                <p>Organize all challenges within each competition for easy reference and prioritization. Filter by category, difficulty, or completion status.</p>
            </div>
            
            <div class="feature-card">
                <div class="feature-title">
                    <div class="feature-icon">📝</div>
                    Collaborative Documentation
                </div>
                <p>Share progress in real-time with a built-in text editor that creates dedicated note documents for each challenge. No more scattered communication.</p>
            </div>
            
            <div class="feature-card">
                <div class="feature-title">
                    <div class="feature-icon">🧠</div>
                    Knowledge Sharing
                </div>
                <p>Ensure all team members can contribute insights and access the latest information about each challenge. Build your team's knowledge base.</p>
            </div>
        </div>
    </div>
</main>

<style>
    .welcome-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    text-align: center;
    padding: 20px;
    }

    logo {
        font-size: 1.5rem;
        font-weight: bold;
    }

    ctf {
        color: #dc4405;
    }
    .collab {
        color: #666;
        margin-left: 2px;
    }

    main h1 {
        font-size: 2.5rem;
        color: #dc4405;
        margin-bottom: 10px;
    }

    main p {
        font-size: 1.2rem;
        color: #666;
    }

    main {
        margin-top: 20px;
    }

    .about-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    text-align: center;
    background-color: #f0f0f0;
    padding: 20px;
    }

    header h1 {
        font-size: 2.5rem;
        color: #333;
        margin-bottom: 10px;
    }

    header p {
        font-size: 1.2rem;
        color: #0c343d;
    }

    main {
        margin-top: 20px;
    }

    footer {
        margin-top: 50px;
        font-size: 0.9rem;
        color: #2f4f4f;
    }
    .form-container {
        display: flex;
        justify-content: center; /* Center the form horizontally */
        width: 100%;
        margin-top: 20px;
    }
    .separator{
        color: #dc4405;
    }
    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    }
    
    body {
        color: white;
        line-height: 1.6;
    }

    .container {
        max-width: 1200px;
        margin: 0 auto;
        padding: 2rem;
    }
    
    h1 {
        font-size: 2.5rem;
        margin-bottom: 1.5rem;
        color: #dc4405;
        text-align: center;
    }
    
    h2 {
        font-size: 1.8rem;
        margin: 2rem 0 1rem;
        color: #dc4405;
    }
    
    p {
        margin-bottom: 1.5rem;
        font-size: 1.1rem;
    }
    
    .features {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
        gap: 2rem;
        margin: 3rem 0;
    }
    
    .feature-card {
        background-color: #1e293b;
        border-radius: 8px;
        padding: 1.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        transition: transform 0.3s ease;
    }
    
    .feature-card:hover {
        transform: translateY(-5px);
    }
    
    .feature-title {
        font-size: 1.4rem;
        color: #dc4405;
        margin-bottom: 1rem;
        display: flex;
        align-items: center;
    }
    
    .feature-icon {
        background-color: #164e63;
        width: 40px;
        height: 40px;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-right: 0.8rem;
    }
            
    footer {
        padding: 2rem;
        text-align: center;
        margin-top: 3rem;
        color: #94a3b8;
    }
</style>

<footer>
    <p1>&copy; 2025 CTF-Collab. All rights reserved.</p1>
</footer>

<!-- Code adapted using Claude AI  -->
  
 

  