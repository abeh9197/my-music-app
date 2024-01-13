// pages/index.tsx
import Link from 'next/link';

const Home = () => {
    return (
        <div>
            <h1>Welcome to My Music App</h1>
            <Link href="/upload">
                Go to Upload Page
            </Link>
        </div>
    );
};

export default Home;
