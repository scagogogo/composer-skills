import React from 'react'
import { useTranslation } from 'react-i18next'
import { Typography } from 'antd'
import { SearchOutlined, CheckCircleOutlined, DownloadOutlined, RocketOutlined } from '@ant-design/icons'
import { motion } from 'framer-motion'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'
import { RevealSection, useScrollReveal } from '../components/ScrollReveal'

const { Paragraph } = Typography

const stepIcons = [
  <SearchOutlined />,
  <CheckCircleOutlined />,
  <DownloadOutlined />,
  <RocketOutlined />,
]

const stepColors = ['#0891B2', '#16A34A', '#2563EB', '#EA580C']

const AutoInstallSection: React.FC = () => {
  const { t } = useTranslation()

  const steps = [
    { title: t('autoInstall.detectTitle'), desc: t('autoInstall.detectDesc') },
    { title: t('autoInstall.checkTitle'), desc: t('autoInstall.checkDesc') },
    { title: t('autoInstall.installTitle'), desc: t('autoInstall.installDesc') },
    { title: t('autoInstall.readyTitle'), desc: t('autoInstall.readyDesc') },
  ]

  const { ref: timelineRef, isVisible } = useScrollReveal({ threshold: 0.1 })

  return (
    <section id="auto-install">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('autoInstall.title')} subtitle={t('autoInstall.subtitle')} />

        {/* Animated timeline */}
        <div ref={timelineRef} style={{ maxWidth: 640, margin: '0 auto 48px' }}>
          {steps.map((step, index) => {
            const color = stepColors[index]
            return (
              <motion.div
                key={index}
                initial={{ opacity: 0, y: 20 }}
                animate={isVisible ? { opacity: 1, y: 0 } : {}}
                transition={{ duration: 0.4, delay: index * 0.12, ease: [0.25, 0.1, 0.25, 1] }}
                style={{
                  display: 'flex',
                  gap: 20,
                  position: 'relative',
                  paddingBottom: index < steps.length - 1 ? 32 : 0,
                }}
              >
                {/* Vertical line */}
                {index < steps.length - 1 && (
                  <div
                    style={{
                      position: 'absolute',
                      left: 19,
                      top: 40,
                      width: 2,
                      height: 'calc(100% - 40px)',
                      background: '#E2E8F0',
                    }}
                  />
                )}

                {/* Step number circle */}
                <div
                  style={{
                    width: 40,
                    height: 40,
                    borderRadius: 4,
                    background: color,
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    color: '#fff',
                    fontSize: 16,
                    flexShrink: 0,
                    position: 'relative',
                    zIndex: 1,
                  }}
                >
                  {stepIcons[index]}
                </div>

                {/* Content */}
                <div style={{ flex: 1, paddingTop: 2 }}>
                  <div style={{ fontWeight: 700, fontSize: 15, color: '#0F172A', marginBottom: 4 }}>
                    {step.title}
                  </div>
                  <Paragraph style={{ color: '#64748B', margin: 0, fontSize: 14, lineHeight: 1.6 }}>
                    {step.desc}
                  </Paragraph>
                </div>
              </motion.div>
            )
          })}
        </div>

        {/* Code example */}
        <RevealSection>
          <div style={{ maxWidth: 640, margin: '0 auto' }}>
            <CodeBlock code={t('autoInstall.code')} />
          </div>
        </RevealSection>
      </div>
    </section>
  )
}

export default AutoInstallSection
