import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function SignIn() {
    return (
        <div className="w-1/4 mx-auto">
            <Card className="rounded-md shadow-none">
                <CardHeader>
                    <CardTitle>
                        Sign In
                    </CardTitle>
                    <CardDescription>
                    Create an account and discover the worlds' best UI component framework.
                    </CardDescription>
                </CardHeader>
                <CardContent className="space-y-2">
                    <div className='space-y-1'>
                        <Label htmlFor="email">E-Mail</Label>
                        <Input placeholder="Enter your E-Mail" className="rounded focus-visible:ring-amber-300 focus-visible:ring-2" />
                    </div>
                    <div className='space-y-1'>
                        <Label htmlFor="email">Password</Label>
                        <Input placeholder="Enter your Password" className="rounded focus-visible:ring-amber-300 focus-visible:ring-2" />
                    </div>
                    <Button variant="amber" className="w-full !mt-4">Log In</Button>
                </CardContent>
            </Card>
        </div>
    )
}