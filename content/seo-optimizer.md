# SEO Optimizer

---

### Summary
> A modern web application for analyzing and optimizing website SEO performance. Built with Go (backend) and TypeScript/React (frontend).

<div style="text-align:center;">
  <img src="/static/images/seo-optimizer.png" alt="website screenshot" style="max-width:60%; height:auto; border-radius:8px; box-shadow:0 4px 12px rgba(0,0,0,0.15);">
</div>

___

### Features:

- Real-time SEO analysis
- Comprehensive metrics including:
  - Title and meta tags optimization
  - Header structure analysis
  - Content quality assessment
  - Performance metrics (load time, page size)
  - Mobile optimization check
  - Internal and external link analysis
- Visual score indicators
- Detailed recommendations for improvement
- Modern, responsive UI
- Enhanced Statistics Dashboard:
  - Monthly statistics tracking with persistence
  - Automatic data retention management
  - Unique visitors tracking
  - Analysis request monitoring
  - Error rate tracking
  - Average load time metrics
  - Private URL tracking (tracked but only visible in development)
  - Cache performance metrics
  - Graceful shutdown with data preservation
  - Production-ready metrics display
  - Development-only detailed URL analysis
- Environment-aware configuration
- Persistent statistics storage with automatic cleanup
- SEO Optimization Features (WIP):

### Skills Used / Developed:
#### Backend
- Go 1.21
- Gin web framework
- goquery for HTML parsing
- Custom statistics tracking
- Automatic monthly data rotation

#### Frontend
- TypeScript / React 18
- Modern CSS
- Real-time statistics updates
- Mobile responsive design

#### Infrastructure
- Docker
- Docker Compose
- Nginx
- Traefik for:
  - SSL termination
  - URL rewriting
  - Secure file serving
- Persistent volumes for data storage
- Private asset management

---

### Links:
- [Github Repo](https://github.com/kingelvyn/seo-optimizer)
- [SEO Optimizer Link](https://seo-optimizer.elvynprise.xyz/)
