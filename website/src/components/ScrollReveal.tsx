import { useEffect, useRef, useState } from 'react'
import { motion, type Variants } from 'framer-motion'

interface UseScrollRevealOptions {
  threshold?: number
  rootMargin?: string
}

export function useScrollReveal(options: UseScrollRevealOptions = {}) {
  const ref = useRef<HTMLDivElement>(null)
  const [isVisible, setIsVisible] = useState(false)

  useEffect(() => {
    const el = ref.current
    if (!el) return

    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsVisible(true)
          observer.unobserve(el)
        }
      },
      {
        threshold: options.threshold ?? 0.15,
        rootMargin: options.rootMargin ?? '0px 0px -40px 0px',
      }
    )

    observer.observe(el)
    return () => observer.disconnect()
  }, [options.threshold, options.rootMargin])

  return { ref, isVisible }
}

export const fadeUpVariants: Variants = {
  hidden: { opacity: 0, y: 30 },
  visible: { opacity: 1, y: 0 },
}

export const fadeInVariants: Variants = {
  hidden: { opacity: 0 },
  visible: { opacity: 1 },
}

export const scaleInVariants: Variants = {
  hidden: { opacity: 0, scale: 0.95 },
  visible: { opacity: 1, scale: 1 },
}

export const staggerContainer: Variants = {
  hidden: {},
  visible: {
    transition: {
      staggerChildren: 0.06,
    },
  },
}

export const staggerItem: Variants = {
  hidden: { opacity: 0, y: 20 },
  visible: { opacity: 1, y: 0, transition: { duration: 0.4 } },
}

interface RevealSectionProps {
  children: React.ReactNode
  className?: string
  style?: React.CSSProperties
  variants?: Variants
  delay?: number
}

export const RevealSection: React.FC<RevealSectionProps> = ({
  children,
  className,
  style,
  variants = fadeUpVariants,
  delay = 0,
}) => {
  const { ref, isVisible } = useScrollReveal()

  return (
    <div ref={ref} className={className} style={style}>
      <motion.div
        initial="hidden"
        animate={isVisible ? 'visible' : 'hidden'}
        variants={variants}
        transition={{ duration: 0.5, delay, ease: [0.25, 0.1, 0.25, 1] }}
      >
        {children}
      </motion.div>
    </div>
  )
}

export const StaggerGrid: React.FC<{
  children: React.ReactNode
  className?: string
  style?: React.CSSProperties
}> = ({ children, className, style }) => {
  const { ref, isVisible } = useScrollReveal()

  return (
    <div ref={ref} className={className} style={style}>
      <motion.div
        initial="hidden"
        animate={isVisible ? 'visible' : 'hidden'}
        variants={staggerContainer}
      >
        {children}
      </motion.div>
    </div>
  )
}

export const StaggerItem: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return <motion.div variants={staggerItem}>{children}</motion.div>
}
