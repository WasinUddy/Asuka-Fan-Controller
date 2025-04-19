import { useEffect, useState } from 'react'

export default function App() {
    const [mode, setMode] = useState('manual')
    const [speed, setSpeed] = useState(50)
    const [loading, setLoading] = useState(false)

    // Round to nearest 5
    const roundToFive = (value : number) => Math.round(value / 5) * 5

    // Preset configurations
    const presets = [
        { name: "Off", value: 0 },
        { name: "Low", value: 10 },
        { name: "Medium", value: 25 },
        { name: "High", value: 50 },
        { name: "Turbo", value: 75 },
        { name: "Berserk", value: 100 }
    ]

    const fetchStatus = async () => {
        try {
            const res = await fetch('/fan/status')
            const data = await res.json()
            setMode(data.mode)
            setSpeed(data.speed)
        } catch (error) {
            console.error("Failed to fetch status:", error)
        }
    }

    useEffect(() => {
        fetchStatus()
    }, [])

    const toggleMode = async () => {
        const newMode = mode === 'auto' ? 'manual' : 'auto'
        setLoading(true)
        try {
            await fetch(`/fan/mode?mode=${newMode}`)
            await fetchStatus()
        } catch (error) {
            console.error("Failed to toggle mode:", error)
        } finally {
            setLoading(false)
        }
    }

    const updateSpeed = async (val :number) => {
        const roundedVal = roundToFive(val)
        setSpeed(roundedVal)
        setLoading(true)
        try {
            await fetch(`/fan/speed?speed=${roundedVal}`)
        } catch (error) {
            console.error("Failed to update speed:", error)
        } finally {
            setLoading(false)
        }
    }

    const handleSliderChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = parseInt(e.target.value)
        setSpeed(value)
    }

    const handleSliderRelease = () => {
        updateSpeed(speed)
    }

    // Custom progress bar for EVA-02 theme
    const ProgressBar = ({ value }: { value : number }) => {
        return (
            <div className="h-3 w-full bg-black rounded-lg overflow-hidden relative">
                <div
                    className="h-full bg-gradient-to-r from-red-600 via-orange-500 to-yellow-500"
                    style={{ width: `${value}%` }}
                />
                {/* EVA-02 style markers */}
                <div className="absolute top-0 bottom-0 left-0 right-0 flex justify-between px-1">
                    {[...Array(10)].map((_, i) => (
                        <div key={i} className="w-0.5 h-full bg-black opacity-30" />
                    ))}
                </div>
            </div>
        )
    }

    return (
        <div className="fixed inset-0 flex items-center justify-center bg-gradient-to-br from-red-800 to-black overflow-auto">
            {/* EVA-02 pattern overlay */}
            <div className="fixed inset-0 opacity-10 pointer-events-none">
                <div className="h-full w-full bg-black">
                    {[...Array(20)].map((_, i) => (
                        <div
                            key={i}
                            className="absolute bg-red-500"
                            style={{
                                height: `${Math.random() * 30 + 5}px`,
                                width: `${Math.random() * 200 + 50}px`,
                                top: `${Math.random() * 100}%`,
                                left: `${Math.random() * 100}%`,
                                transform: `rotate(${Math.random() * 360}deg)`,
                                opacity: 0.3
                            }}
                        />
                    ))}
                </div>
            </div>

            <div className="w-full max-w-xl p-6 rounded-lg shadow-xl border-2 border-orange-500 bg-gradient-to-b from-red-900 to-black relative overflow-hidden z-10 m-4">
                {/* Smaller Angular EVA-style decorative elements */}
                <div className="absolute top-0 left-0 w-16 h-16 bg-red-600 opacity-40" style={{ clipPath: 'polygon(0 0, 0% 100%, 100% 0)' }} />
                <div className="absolute top-0 right-0 w-16 h-16 bg-orange-500 opacity-40" style={{ clipPath: 'polygon(100% 0, 0 0, 100% 100%)' }} />
                <div className="absolute bottom-0 left-0 w-16 h-16 bg-yellow-500 opacity-40" style={{ clipPath: 'polygon(0 100%, 100% 100%, 0 0)' }} />
                <div className="absolute bottom-0 right-0 w-16 h-16 bg-red-800 opacity-40" style={{ clipPath: 'polygon(100% 100%, 0 100%, 100% 0)' }} />

                {/* Header with EVA-02 inspired styling */}
                <div className="relative mb-8 border-b-2 border-orange-500 pb-4">
                    <div className="flex items-center justify-center">
                        <h1 className="text-3xl font-bold">
                            <span className="text-red-500">Asuka</span>
                            <span className="text-white">-</span>
                            <span className="text-red-500">Chan</span>
                            <span className="text-yellow-400"> Fan Control</span>
                        </h1>
                    </div>
                    <div className="text-center text-orange-300 text-sm mt-1">TrueNAS Dell PowerEdge R230</div>
                </div>

                {/* Mode Toggle */}
                <div className="mb-8">
                    <div className="flex justify-between items-center mb-3">
                        <span className="text-white font-medium">OPERATION MODE:</span>
                        <span className="font-bold text-yellow-400">{mode.toUpperCase()}</span>
                    </div>
                    <button
                        onClick={toggleMode}
                        disabled={loading}
                        className="w-full bg-gradient-to-r from-red-700 to-red-900 hover:from-red-600 hover:to-red-800 disabled:from-gray-700 disabled:to-gray-900 text-white py-3 rounded transition-colors border border-orange-500 shadow-md font-bold"
                    >
                        {loading ? "PROCESSING..." : mode === 'auto' ? 'SWITCH TO MANUAL CONTROL' : 'SWITCH TO AUTO CONTROL'}
                    </button>
                </div>

                {/* Fan Speed Control - Only visible in manual mode */}
                {mode === 'manual' && (
                    <div className="mb-8">
                        <div className="flex justify-between items-center mb-3">
                            <span className="text-white font-medium">FAN SPEED:</span>
                            <div className="flex items-center">
                                <span className="text-xl font-bold text-yellow-400">{speed}</span>
                                <span className="text-yellow-400 ml-1">%</span>
                            </div>
                        </div>

                        {/* Custom styled slider with proper positioning */}
                        <div className="mb-4 relative">
                            <ProgressBar value={speed} />
                            <input
                                type="range"
                                min="0"
                                max="100"
                                step="5"
                                value={speed}
                                onChange={handleSliderChange}
                                onMouseUp={handleSliderRelease}
                                onTouchEnd={handleSliderRelease}
                                disabled={loading}
                                className="absolute inset-0 w-full h-6 opacity-0 cursor-pointer"
                            />
                        </div>

                        {/* EVA-style preset buttons */}
                        <div className="grid grid-cols-6 gap-2 mt-6">
                            {presets.map((preset) => (
                                <button
                                    key={preset.name}
                                    onClick={() => updateSpeed(preset.value)}
                                    disabled={loading}
                                    className={`py-2 px-1 ${
                                        speed === preset.value
                                            ? 'bg-gradient-to-r from-red-600 to-red-800 border-yellow-400'
                                            : 'bg-gradient-to-r from-gray-900 to-black border-orange-700 hover:border-orange-500'
                                    } text-white rounded text-xs transition-all font-medium border`}
                                >
                                    {preset.name.toUpperCase()}
                                </button>
                            ))}
                        </div>
                    </div>
                )}
            </div>
        </div>
    )
}