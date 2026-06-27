import React, { useMemo, Suspense, lazy } from 'react'
import { useTranslation } from 'react-i18next'
import { Typography, Collapse, Spin } from 'antd'
import SectionTitle from '../components/SectionTitle'
import { RevealSection } from '../components/ScrollReveal'

const { Title } = Typography

// Lazy-load the chart component to reduce initial bundle size
const ColumnChart = lazy(() =>
  import('@ant-design/charts').then((mod) => ({ default: mod.Column }))
)

const CoverageSection: React.FC = () => {
  const { t } = useTranslation()

  const composerCategories = t('coverage.composerCategories', { returnObjects: true }) as Array<{
    category: string
    count: string
    highlights: string
  }>

  const packagistCategories = t('coverage.packagistCategories', { returnObjects: true }) as Array<{
    category: string
    methods: string
  }>

  // Prepare chart data from composer categories
  const chartData = useMemo(() => {
    return composerCategories.map((cat) => ({
      category: cat.category,
      count: parseInt(cat.count, 10),
    }))
  }, [composerCategories])

  const chartConfig = useMemo(
    () => ({
      data: chartData,
      xField: 'category',
      yField: 'count',
      color: '#2563EB',
      style: {
        borderRadius: 2,
      },
      axis: {
        x: {
          label: {
            style: { fontSize: 11, fill: '#64748B' },
            autoRotate: true,
            autoHide: false,
          },
          line: { style: { stroke: '#E2E8F0' } },
          tickLine: null,
        },
        y: {
          label: {
            style: { fontSize: 11, fill: '#64748B' },
          },
          grid: { line: { style: { stroke: '#F1F5F9', lineWidth: 1 } } },
          line: null,
          tickLine: null,
        },
      },
      label: {
        text: (d: { count: number }) => String(d.count),
        textBaseline: 'bottom',
        style: { fontSize: 11, fontWeight: 600, fill: '#334155' },
      },
      tooltip: {
        title: 'category',
        items: [{ channel: 'y', name: 'Methods' }],
      },
      height: 280,
    }),
    [chartData]
  )

  // Packagist table
  const packagistRows = packagistCategories.map((item, i) => (
    <tr key={i} style={{ borderBottom: '1px solid #F1F5F9' }}>
      <td style={{ padding: '8px 12px', fontWeight: 600, fontSize: 13, width: '25%', color: '#0F172A' }}>
        {item.category}
      </td>
      <td style={{ padding: '8px 12px', fontSize: 13, color: '#475569' }}>
        {item.methods}
      </td>
    </tr>
  ))

  // Composer table
  const composerRows = composerCategories.map((item, i) => (
    <tr key={i} style={{ borderBottom: '1px solid #F1F5F9' }}>
      <td style={{ padding: '8px 12px', fontWeight: 600, fontSize: 13, width: '20%', color: '#0F172A' }}>
        {item.category}
      </td>
      <td style={{ padding: '8px 12px', width: '12%', textAlign: 'center' }}>
        <span
          style={{
            background: '#EFF6FF',
            color: '#2563EB',
            padding: '1px 8px',
            borderRadius: 2,
            fontWeight: 600,
            fontSize: 13,
          }}
        >
          {item.count}
        </span>
      </td>
      <td style={{ padding: '8px 12px', fontSize: 13, color: '#475569' }}>
        {item.highlights}
      </td>
    </tr>
  ))

  const collapseItems = [
    {
      key: 'packagist',
      label: <Title level={4} style={{ margin: 0, fontSize: 15 }}>{t('coverage.packagistTitle')}</Title>,
      children: (
        <div style={{ overflowX: 'auto' }}>
          <table style={{ width: '100%', borderCollapse: 'collapse' }}>
            <thead>
              <tr style={{ background: '#F8FAFC', borderBottom: '2px solid #E2E8F0' }}>
                <th style={{ padding: '8px 12px', textAlign: 'left', fontSize: 13, fontWeight: 600, color: '#475569' }}>
                  {t('coverage.categoryHeader')}
                </th>
                <th style={{ padding: '8px 12px', textAlign: 'left', fontSize: 13, fontWeight: 600, color: '#475569' }}>
                  {t('coverage.methodsHeader')}
                </th>
              </tr>
            </thead>
            <tbody>{packagistRows}</tbody>
          </table>
        </div>
      ),
    },
    {
      key: 'composer',
      label: <Title level={4} style={{ margin: 0, fontSize: 15 }}>{t('coverage.composerTitle')}</Title>,
      children: (
        <div style={{ overflowX: 'auto' }}>
          <table style={{ width: '100%', borderCollapse: 'collapse' }}>
            <thead>
              <tr style={{ background: '#F8FAFC', borderBottom: '2px solid #E2E8F0' }}>
                <th style={{ padding: '8px 12px', textAlign: 'left', fontSize: 13, fontWeight: 600, color: '#475569' }}>
                  {t('coverage.categoryHeader')}
                </th>
                <th style={{ padding: '8px 12px', textAlign: 'center', fontSize: 13, fontWeight: 600, color: '#475569' }}>
                  {t('coverage.countHeader')}
                </th>
                <th style={{ padding: '8px 12px', textAlign: 'left', fontSize: 13, fontWeight: 600, color: '#475569' }}>
                  {t('coverage.highlightsHeader')}
                </th>
              </tr>
            </thead>
            <tbody>{composerRows}</tbody>
          </table>
        </div>
      ),
    },
  ]

  return (
    <section id="coverage" className="section-alt">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('coverage.title')} subtitle={t('coverage.subtitle')} />

        {/* Bar chart — method counts by category */}
        <RevealSection>
          <div style={{ maxWidth: 860, margin: '0 auto 48px' }}>
            <div
              style={{
                border: '1px solid #E2E8F0',
                background: '#fff',
                padding: '20px 20px 12px',
              }}
            >
              <Title level={5} style={{ marginBottom: 8, fontSize: 14, color: '#475569', fontWeight: 600 }}>
                {t('coverage.composerTitle')}
              </Title>
              <Suspense fallback={<div style={{ textAlign: 'center', padding: 40 }}><Spin /></div>}>
                <ColumnChart {...chartConfig} />
              </Suspense>
            </div>
          </div>
        </RevealSection>

        {/* Detail tables */}
        <Collapse
          defaultActiveKey={['packagist', 'composer']}
          items={collapseItems}
          style={{ background: '#fff' }}
        />
      </div>
    </section>
  )
}

export default CoverageSection
