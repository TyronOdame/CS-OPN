'use client';

import { useState } from 'react';

export default function LoginPage() {
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [currentView, setCurrentView] = useState<
    'login' | 'register' | 'forgot'
  >('login');

  return (
    <div className="min-h-screen flex font-sans bg-zinc-950">
      {/* Left Panel - CS:OPN Branding */}
      <div className="hidden lg:flex lg:w-1/2 relative overflow-hidden bg-[#0f0f0f]">
        <div className="relative z-10 flex flex-col justify-between w-full px-12 py-12">
          <div className="flex items-center">
            <div className="w-10 h-10 bg-[#ea580c] rounded-lg flex items-center justify-center mr-3">
              <span className="text-black font-bold text-lg">CS</span>
            </div>
            <h1 className="text-xl font-semibold text-white">CS:OPN</h1>
          </div>

          <div className="flex-1 flex flex-col justify-center">
            <h2 className="text-4xl text-white mb-6 leading-tight">
              Experience the thrill of case opening without spending real money.
            </h2>
            <p className="text-white/90 text-lg leading-relaxed">
              Log in to access your inventory, open cases, and trade with
              friends.
            </p>
          </div>

          <div className="flex justify-between items-center text-white/70 text-sm">
            <span>Copyright © 2025 CS:OPN</span>
            <span className="cursor-pointer hover:text-white/90">
              Privacy Policy
            </span>
          </div>
        </div>
      </div>

      {/* Right Panel - Login Form */}
      <div className="w-full lg:w-1/2 flex items-center justify-center p-8 bg-zinc-950">
        <div className="w-full max-w-md space-y-8">
          {/* Mobile Logo */}
          <div className="lg:hidden text-center mb-8">
            <div className="w-10 h-10 bg-[#ea580c] rounded-lg flex items-center justify-center mx-auto mb-3">
              <span className="text-white font-bold text-lg">CS</span>
            </div>
            <h1 className="text-xl font-semibold text-white">CS:OPN</h1>
          </div>

          <div className="space-y-6">
            <div className="space-y-2 text-center">
              {currentView === 'forgot' && (
                <button
                  onClick={() => setCurrentView('login')}
                  className="absolute left-8 top-8 p-2 hover:bg-zinc-800 rounded-md cursor-pointer text-white"
                >
                  <svg
                    className="h-4 w-4"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M15 19l-7-7 7-7"
                    />
                  </svg>
                </button>
              )}
              <h2 className="text-3xl text-white">
                {currentView === 'login' && 'Welcome Back'}
                {currentView === 'register' && 'Create Account'}
                {currentView === 'forgot' && 'Reset Password'}
              </h2>
              <p className="text-zinc-400">
                {currentView === 'login' &&
                  'Enter your email and password to access your account.'}
                {currentView === 'register' &&
                  'Create a new account to get started with CS:OPN.'}
                {currentView === 'forgot' &&
                  "Enter your email address and we'll send you a reset link."}
              </p>
            </div>

            <div className="space-y-4">
              {currentView === 'register' && (
                <div className="space-y-2">
                  <label
                    htmlFor="name"
                    className="text-sm font-medium text-white block"
                  >
                    Full Name
                  </label>
                  <input
                    id="name"
                    type="text"
                    placeholder="John Doe"
                    className="w-full h-12 px-4 border border-zinc-800 focus:ring-0 focus:outline-none shadow-none rounded-lg bg-zinc-900 focus:border-[#ea580c] text-white placeholder:text-zinc-500"
                  />
                </div>
              )}

              <div className="space-y-2">
                <label
                  htmlFor="email"
                  className="text-sm font-medium text-white block"
                >
                  Email
                </label>
                <input
                  id="email"
                  type="email"
                  placeholder="user@example.com"
                  className="w-full h-12 px-4 border border-zinc-800 focus:ring-0 focus:outline-none shadow-none rounded-lg bg-zinc-900 focus:border-[#ea580c] text-white placeholder:text-zinc-500"
                />
              </div>

              {currentView !== 'forgot' && (
                <div className="space-y-2">
                  <label
                    htmlFor="password"
                    className="text-sm font-medium text-white block"
                  >
                    Password
                  </label>
                  <div className="relative">
                    <input
                      id="password"
                      type={showPassword ? 'text' : 'password'}
                      placeholder="Enter password"
                      className="w-full h-12 px-4 pr-10 border border-zinc-800 focus:ring-0 focus:outline-none shadow-none rounded-lg bg-zinc-900 focus:border-[#ea580c] text-white placeholder:text-zinc-500"
                    />
                    <button
                      type="button"
                      className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent cursor-pointer text-zinc-400"
                      onClick={() => setShowPassword(!showPassword)}
                    >
                      {showPassword ? (
                        <svg
                          className="h-4 w-4"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"
                          />
                        </svg>
                      ) : (
                        <svg
                          className="h-4 w-4"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                          />
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                          />
                        </svg>
                      )}
                    </button>
                  </div>
                </div>
              )}

              {currentView === 'register' && (
                <div className="space-y-2">
                  <label
                    htmlFor="confirmPassword"
                    className="text-sm font-medium text-white block"
                  >
                    Confirm Password
                  </label>
                  <div className="relative">
                    <input
                      id="confirmPassword"
                      type={showConfirmPassword ? 'text' : 'password'}
                      placeholder="Confirm password"
                      className="w-full h-12 px-4 pr-10 border border-zinc-800 focus:ring-0 focus:outline-none shadow-none rounded-lg bg-zinc-900 focus:border-[#ea580c] text-white placeholder:text-zinc-500"
                    />
                    <button
                      type="button"
                      className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent cursor-pointer text-zinc-400"
                      onClick={() =>
                        setShowConfirmPassword(!showConfirmPassword)
                      }
                    >
                      {showConfirmPassword ? (
                        <svg
                          className="h-4 w-4"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"
                          />
                        </svg>
                      ) : (
                        <svg
                          className="h-4 w-4"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                          />
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                          />
                        </svg>
                      )}
                    </button>
                  </div>
                </div>
              )}

              {currentView === 'login' && (
                <div className="flex items-center justify-end">
                  <button
                    className="p-0 h-auto text-sm text-[#ea580c] hover:text-[#ea580c]/80 cursor-pointer underline-offset-4 hover:underline"
                    onClick={() => setCurrentView('forgot')}
                  >
                    Forgot Your Password?
                  </button>
                </div>
              )}
            </div>

            <button className="w-full h-12 text-sm font-medium bg-[#ea580c] text-white hover:bg-[#ea580c]/90 rounded-lg shadow-none cursor-pointer transition-colors">
              {currentView === 'login' && 'Log In'}
              {currentView === 'register' && 'Create Account'}
              {currentView === 'forgot' && 'Send Reset Link'}
            </button>

            <div className="text-center text-sm text-zinc-400">
              {currentView === 'login' && (
                <>
                  Don't Have An Account?{' '}
                  <button
                    className="p-0 h-auto text-sm text-[#ea580c] hover:text-[#ea580c]/80 font-medium cursor-pointer underline-offset-4 hover:underline"
                    onClick={() => setCurrentView('register')}
                  >
                    Register Now.
                  </button>
                </>
              )}
              {currentView === 'register' && (
                <>
                  Already Have An Account?{' '}
                  <button
                    className="p-0 h-auto text-sm text-[#ea580c] hover:text-[#ea580c]/80 font-medium cursor-pointer underline-offset-4 hover:underline"
                    onClick={() => setCurrentView('login')}
                  >
                    Sign In.
                  </button>
                </>
              )}
              {currentView === 'forgot' && (
                <>
                  Remember Your Password?{' '}
                  <button
                    className="p-0 h-auto text-sm text-[#ea580c] hover:text-[#ea580c]/80 font-medium cursor-pointer underline-offset-4 hover:underline"
                    onClick={() => setCurrentView('login')}
                  >
                    Back to Login.
                  </button>
                </>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
