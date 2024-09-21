import Navbar from "@/components/layouts/navbar";
import { css } from "styled-system/css";

export default function LandingLayout({ children }: {children: React.ReactNode}) {
    return (
        <div className={css({
            minH: 'screen'
        })}>
            <Navbar />
            {children}
        </div>
    )
}