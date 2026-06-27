import React from 'react'
import { useTranslation } from 'react-i18next'
import { Typography } from 'antd'
import { BookOutlined, CodeOutlined, CloudOutlined, CodeSandboxOutlined, BuildOutlined } from '@ant-design/icons'
import { motion } from 'framer-motion'
import SectionTitle from '../components/SectionTitle'
import { StaggerGrid, StaggerItem, useScrollReveal } from '../components/ScrollReveal'

const { Paragraph } = Typography

const layerIcons = [
  <BookOutlined />,
  <CodeSandboxOutlined />,
  <CloudOutlined />,
  <CodeOutlined />,
  <BuildOutlined />,
]

const layerColors = [
  { bg: '#EFF6FF', border: '#BFDBFE', text: '#1D4ED8', iconBg: '#2563EB' },
  { bg: '#F0FDF4', border: '#BBF7D0', text: '#15803D', iconBg: '#16A34A' },
  { bg: '#ECFEFF', border: '#A5F3FC', text: '#0E7490', iconBg: '#0891B2' },
  { bg: '#FFF7ED', border: '#FED7AA', text: '#C2410C', iconBg: '#EA580C' },
  { bg: '#F8FAFC', border: '#E2E8F0', text: '#334155', iconBg: '#475569' },
]

const ArchitectureSection: React.FC = () => {
  const { t } = useTranslation()

  const layers = t('architecture.layers', { returnObjects: true }) as Array<{
    layer: string
    func: string
    pkg: string
  }>

  const { ref: layerRef, isVisible } = useScrollReveal({ threshold: 0.1 })

  return (
    <section id="architecture" className="section-alt">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('architecture.title')} subtitle={t('architecture.subtitle')} />

        {/* Layered architecture diagram */}
        <div ref={layerRef} style={{ maxWidth: 720, margin: '0 auto 48px' }}>
          {layers.map((layer, index) => {
            const color = layerColors[index]
            return (
              <motion.div
                key={index}
                initial={{ opacity: 0, x: -30 }}
                animate={isVisible ? { opacity: 1, x: 0 } : {}}
                transition={{ duration: 0.45, delay: index * 0.1, ease: [0.25, 0.1, 0.25, 1] }}
              >
                <div
                  style={{
                    display: 'flex',
                    alignItems: 'stretch',
                    marginBottom: index < layers.length - 1 ? 8 : 0,
                    position: 'relative',
                  }}
                >
                  {/* Arrow connector */}
                  {index < layers.length - 1 && (
                    <div
                      style={{
                        position: 'absolute',
                        bottom: -8,
                        left: '50%',
                        transform: 'translateX(-50%)',
                        width: 2,
                        height: 8,
                        background: '#CBD5E1',
                      }}
                    />
                  )}

                  {/* Icon column */}
                  <div
                    style={{
                      width: 44,
                      minHeight: 64,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      background: color.iconBg,
                      border: `1px solid ${color.border}`,
                      borderRight: 'none',
                      color: '#fff',
                      fontSize: 18,
                      flexShrink: 0,
                    }}
                  >
                    {layerIcons[index]}
                  </div>

                  {/* Content */}
                  <div
                    style={{
                      flex: 1,
                      background: color.bg,
                      border: `1px solid ${color.border}`,
                      borderLeft: 'none',
                      padding: '12px 16px',
                      display: 'flex',
                      alignItems: 'center',
                      gap: 16,
                      flexWrap: 'wrap',
                    }}
                  >
                    <div style={{ minWidth: 140 }}>
                      <div style={{ fontWeight: 700, fontSize: 14, color: color.text }}>
                        {layer.layer}
                      </div>
                    </div>
                    <div style={{ flex: 1, minWidth: 160 }}>
                      <Paragraph style={{ margin: 0, fontSize: 13, color: '#475569', lineHeight: 1.5 }}>
                        {layer.func}
                      </Paragraph>
                    </div>
                    <div>
                      <code
                        style={{
                          background: '#fff',
                          padding: '2px 8px',
                          borderRadius: 2,
                          fontSize: 12,
                          color: color.text,
                          border: `1px solid ${color.border}`,
                          fontWeight: 500,
                        }}
                      >
                        {layer.pkg}
                      </code>
                    </div>
                  </div>
                </div>
              </motion.div>
            )
          })}
        </div>

        {/* Summary cards */}
        <StaggerGrid style={{ maxWidth: 720, margin: '0 auto' }}>
          <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(150px, 1fr))', gap: 12 }}>
            {[
              { label: 'Layers', value: '5', color: '#2563EB' },
              { label: 'Packages', value: '6+', color: '#16A34A' },
              { label: 'SDK Methods', value: '254', color: '#0891B2' },
              { label: 'CLI Commands', value: '50+', color: '#EA580C' },
            ].map((stat, i) => (
              <StaggerItem key={i}>
                <div
                  style={{
                    textAlign: 'center',
                    border: `1px solid ${stat.color}30`,
                    background: '#fff',
                    padding: '16px 12px',
                  }}
                >
                  <div style={{ fontSize: 24, fontWeight: 800, color: stat.color, lineHeight: 1.2 }}>
                    {stat.value}
                  </div>
                  <div style={{ fontSize: 12, color: '#64748B', marginTop: 4, fontWeight: 500 }}>
                    {stat.label}
                  </div>
                </div>
              </StaggerItem>
            ))}
          </div>
        </StaggerGrid>
      </div>
    </section>
  )
}

export default ArchitectureSection
