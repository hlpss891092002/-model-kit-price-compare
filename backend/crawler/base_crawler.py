import aiohttp
import asyncio
from bs4 import BeautifulSoup
from typing import List, Dict, Optional
from datetime import datetime
import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class BaseCrawler:
    def __init__(self, store_name: str):
        self.store_name = store_name
        self.session = None
        self.headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36'
        }

    async def __aenter__(self):
        self.session = aiohttp.ClientSession(headers=self.headers)
        return self

    async def __aexit__(self, exc_type, exc_val, exc_tb):
        if self.session:
            await self.session.close()

    async def fetch_page(self, url: str) -> Optional[str]:
        try:
            async with self.session.get(url) as response:
                if response.status == 200:
                    return await response.text()
                else:
                    logger.error(f"Failed to fetch {url}: Status {response.status}")
                    return None
        except Exception as e:
            logger.error(f"Error fetching {url}: {str(e)}")
            return None

    async def parse_product(self, html: str, url: str) -> Dict:
        raise NotImplementedError("Must implement parse_product method")

    async def search_products(self, keyword: str) -> List[Dict]:
        raise NotImplementedError("Must implement search_products method")

    def calculate_total_price(self, price: float, shipping: float) -> float:
        return price + shipping

    def create_product_dict(
        self,
        name: str,
        price: float,
        shipping: float,
        url: str,
        image_url: Optional[str] = None
    ) -> Dict:
        return {
            "name": name,
            "price": price,
            "shipping_cost": shipping,
            "total_price": self.calculate_total_price(price, shipping),
            "store": self.store_name,
            "url": url,
            "image_url": image_url,
            "last_updated": datetime.now().isoformat()
        }