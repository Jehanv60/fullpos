import { redirect } from '@sveltejs/kit';

// daftar route publik yang bisa diakses tanpa login
const publicPaths = ['/'];

export async function handle({ event, resolve }) {
    const token = event.cookies.get('token');
    const path = event.url.pathname;

    // cek apakah route termasuk route publik
    const isProtected = publicPaths.some(publicPath =>
        path === publicPath || path.startsWith(publicPath + '/')
    );

    if (!isProtected && !token) {
        throw redirect(302, '/'); // redirect ke login
    }
    if (isProtected && token) {
        throw redirect(302, '/dashboard'); // redirect ke login
    }

    return resolve(event);
}
