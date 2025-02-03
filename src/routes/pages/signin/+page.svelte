<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";

</script>

<div class="login-container">
    <header>
        <h1>Login as Administrator</h1>
        <p>Enter your email below to access existing admin account</p>
        
        <div class="form-container">
            <form class="flex flex-col w-full max-w-sm items-center space-y-2">
                <Input type="email" placeholder="me@example.com..." />
                <Input type="password" placeholder="Enter your password" />
                <Button href="src\routes\admin-account.svelte" type="submit">Log In</Button>
            </form>
        </div>
        
        
    </header>

    <main>
    </main>

    <footer>
        <p>&copy; 2025 CTF-Collab. All rights reserved.</p>
    </footer>
</div>

<style>
    .login-container {
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
        color: #dc4405;
    }

    main {
        margin-top: 20px;
    }

    footer {
        margin-top: 40px;
        font-size: 0.9rem;
        color: #777;
    }
    .form-container {
        display: flex;
        justify-content: center; /* Center the form horizontally */
        width: 100%;
        margin-top: 20px;
    }
</style>